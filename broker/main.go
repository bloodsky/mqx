package main

import (
	"fmt"
	"google.golang.org/grpc"
	addservice "mqx/proto"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	AT_LEAST_ONCE = 499
	TIMEOUT_BASED = 500
)

// "Sport" sliced-queue with unique name
var Queue = make([]Message, 0)
// Sliced-list of subscribers for "Sport queue"
var Subscribers = make([]string,0)
// Type of semantic
var Semantic int
// Retrasmit timeout
var re_timeout time.Duration
// Visibility Timeout
var vi_timeout time.Duration
// Channel to make exclusive access --> one single message for one single subscriber per time
var ch int

func main() {
	// Opening a TCP listner
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()
	ch = 0
	// Semantic type to use for the queue
	if strings.Compare(os.Args[1],"ALO" ) == 0 {
		Semantic = AT_LEAST_ONCE // --> Request Retrasmit
		re_timeout, _ = ParseTime(os.Args[2])
	} else if strings.Compare(os.Args[1], "TOB") == 0 {
		Semantic = TIMEOUT_BASED
		vi_timeout, _ = ParseTime(os.Args[2])
	}

	addServ := AddServiceServer{}
	grpcServer := grpc.NewServer()
	addservice.RegisterAddServiceServer(grpcServer, &addServ)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
	}
}

func ParseTime(t string) (time.Duration, error) {
	var mins, hours int
	var err error

	parts := strings.SplitN(t, ":", 2)

	switch len(parts) {
	case 1:
		mins, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
	case 2:
		hours, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}

		mins, err = strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}
	default:
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	if mins > 59 || mins < 0 || hours > 23 || hours < 0 {
		return 0, fmt.Errorf("invalid time: %s", t)
	}

	return time.Duration(hours)*time.Hour + time.Duration(mins)*time.Minute, nil
}