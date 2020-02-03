package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	addservice "mqx/proto"
	"os"
	"strings"
	"time"
)

var StoredMessage = make([]string, 0)

func main() {

	conn, err := grpc.Dial(
		"localhost:1111", grpc.WithInsecure(),)
	if err != nil {
		fmt.Println(err)
	}

	defer conn.Close()
	addServ := addservice.NewAddServiceClient(conn)


	if strings.Compare(os.Args[1],"pub" ) == 0 {
		// user input
		var text string
		fmt.Println("Welcome publisher!\nWrite some message for the queue or the command 'list'")

		for {
			_, _ = fmt.Scan(&text)

			// Sending time
			t := time.Now()
			t.Format("20060102150405")

			// Security bug?
			if strings.Compare(text,"list") == 0 {
				response, err := addServ.List(context.Background(), &addservice.ListQueuedMessage{Text:"list"})
				if err != nil {
					fmt.Println(err)
				} else {
					result := response.List
					fmt.Println(result)
				}
			} else {
				response, err := addServ.Pub(context.Background(), &addservice.TopicMessage{SentOn: t.String(), Text: text})
				if err != nil {
					fmt.Println(err)
				} else {
					result := response.Ack
					if result {
						fmt.Println("Succesfully delivered message to the broker!")
					}
				}
			}
		}
	} else if strings.Compare(os.Args[1],"sub" ) == 0 {
		fmt.Println("Welcome subscriber! Insert your name:")
		var text string

		_, _ = fmt.Scan(&text)

		Topic := "Sport"

		response, err := addServ.Sub(context.Background(), &addservice.TopicSubscription{Name: text, Topic: Topic})

		if err != nil {
			fmt.Println(err)
		} else {
			result := response.Ack
			if result {
				fmt.Println("Succesfully subscribed to 'Sport' topic!")
			}

			for {
				time.Sleep(400*time.Millisecond) // slowing down subscriber
				response, err := addServ.Serve(context.Background(), &addservice.ReadyMessage{Ready:1}) //ready!
				if err != nil {
					fmt.Println(err)
				} else {
					result := response.Text
					var txt string
					if strings.Compare(result,"NOMSG") == 0 {
						fmt.Println("No more work to do, now wait for some publisher to do his work! 'check' command")
						_, _ = fmt.Scan(&txt)
					} else {
						if len(StoredMessage) == 0 {
							StoredMessage = append(StoredMessage, result)
						} else {
							for _, ele := range StoredMessage {
								if ele == result {
									fmt.Println(".. Triggered RR on --> ",result)
								}
							}
							StoredMessage = append(StoredMessage, result)
						}
						fmt.Println(".. delivered message --> ",result)
					}
				}
			}
		}
	}
}


/*func IDgenerator() int32 {
	rand.Seed(time.Now().UnixNano())
	min := 1045
	max := 3045678
	return int32((rand.Intn(max-min+1) + min))
}*/