# MQX
Exercise for "Distributed System" course at University of Tor Vergata.
Simple message queue written in Go for supporting asynchronous comunication between distributed components.

## Developed by: [Pavia Roberto](https://github.com/bloodsky)

# Usage

Unpack everything inside your local Go directory under the src folder.

**Broker**
```
cd broker
go run *.go semantic timeout
```
**Pub/Sub**
```
cd client
go run *.go pub/sub
```
# Example
```
go run *.go ALO 30 
```
- ALO means "AtLeastOnce" and in this example with 30 milliseconds of retrasmit timeout
```
go run *.go TOB 50 
```
- TOB means "Timeout based" and in this example with 50 milliseconds of visibility timeout
```
go run *.go pub 
```
- Publisher can request a "List" operation to get the single message state of the entire queue -> list
```
go run *.go sub 
```
- Subscriber can "Check" for instantaneous update also manually -> check


# Test

* => Tested with a sleep to slow down subscriber

First 2 colomun => ALO with 500ms timeout
Last 2 column => ALO with 100ms timeout 

 319 messages* | 3000 messages | 31743 messages | 49890 messages
------------ | ------------- | ------------- | -------------
Delivered in 02:03.02 | Delivered in 00:04.92 | Delivered in 00:18.22 | Delivered in 00:48.01
Delivered in 02:01.09 | Delivered in 00:03.25 | Delivered in 00:14.92 | Delivered in 00:51.72
Delivered in 02:00.47 | Delivered in 00:03.72 | Delivered in 00:15.48 | Delivered in 00:47.22

# Current bugs

- There are some problems with logging information server side. Multiple write in the log are in a state of overhanging.
- In some particular cases, subscriber need to check for message in queue after a "long" run
- 
