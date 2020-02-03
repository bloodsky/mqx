package main

import (
	"context"
	"fmt"
	addservice "mqx/proto"
	"time"
)

type AddServiceServer struct {
}

func (*AddServiceServer) Serve(ctx context.Context, in *addservice.ReadyMessage) (*addservice.Message, error) {

	if ch == 0 {
		ch = 1 // take a channel spot
		if Semantic == AT_LEAST_ONCE {

			if len(Queue) != 0 {
				tosent := Queue[0].Text

				ctx, cancel := context.WithTimeout(context.Background(), re_timeout*time.Millisecond)
				defer cancel()

				select {
				case <-time.After(re_timeout * time.Millisecond):
					fmt.Println("overslept")
					ch = 0
					return &addservice.Message{Text: tosent}, nil
				case <-ctx.Done():
					Queue = Queue[1:]      // removing message from the queue ONLY if broker received the ack from sub!
					fmt.Println(ctx.Err()) // prints "context deadline exceeded"
					ch = 0
					return &addservice.Message{Text: tosent}, nil
				}
			} else {
				ch = 0
				return &addservice.Message{Text: "NOMSG"}, nil
			}
		} else { // TIMEOUT BASED
			if len(Queue) != 0 {
				tosent := Queue[0].Text
				Queue[0].Visibility = false

				ctx, cancel := context.WithTimeout(context.Background(), vi_timeout*time.Millisecond)
				defer cancel()

				select {
				case <-time.After(re_timeout * time.Millisecond):
					Queue[0].Visibility = false // set visibility message to false
					fmt.Println("overslept")
					ch = 0
					return &addservice.Message{Text: tosent}, nil
				case <-ctx.Done():
					Queue = Queue[1:]      // removing message from the queue ONLY if broker received the ack from sub!
					fmt.Println(ctx.Err())
					ch = 0
					return &addservice.Message{Text: tosent}, nil
				}

			} else {
				ch = 0
				return &addservice.Message{Text: "NOMSG"}, nil
			}
		}
	} else {
		//more consumer
	}
	return &addservice.Message{Text: "NOMSG"}, nil
}


func (*AddServiceServer) List(ctx context.Context, in *addservice.ListQueuedMessage) (*addservice.ReceivedQueuedMessageList, error) {

	if len(Queue) != 0 {
		m := make(map[string]bool)
		for _, ele := range Queue {
			m[ele.Text] = ele.Visibility
		}
		return &addservice.ReceivedQueuedMessageList{List:m}, nil
	} else {
		m := make(map[string]bool)
		m["zero"] = true
		return &addservice.ReceivedQueuedMessageList{List:m}, nil
	}
}


func (*AddServiceServer) Sub(ctx context.Context, in *addservice.TopicSubscription) (*addservice.SubscriberDeliveredMessage, error) {

	// Filling sliced-list of subscribers
	if len(Subscribers) == 0 {
		Subscribers = append(Subscribers, in.Name)
	} else {
		for _, ele := range Subscribers {
			if ele == in.Name {
				return &addservice.SubscriberDeliveredMessage{Ack: false}, nil
			}
		}
	}
	return &addservice.SubscriberDeliveredMessage{Ack: true}, nil
	// Testing
	// log.Print("Registered '",Subscribers[0],"' to Sport queue")
}


func (*AddServiceServer) Pub(ctx context.Context, in *addservice.TopicMessage) (*addservice.QueuedMessage, error) {

	// Filling sliced-queue with incoming message
	msg := NewMessage(in.Text,true)
	Queue = append(Queue, *msg)

	// testing
	// log.Print("got message:",Queue[0])
	return &addservice.QueuedMessage{Ack: true}, nil
}
