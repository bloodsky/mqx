# mqx
Exercise for "Distributed System" course at University of Tor Vergata.
Simple message queue written in Go for supporting asynchronous comunication between distributed components.

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
- ALO -> AtLeastOnce with 30 milliseconds of timeout
```
go run *.go TOB 50 
```
- TOB -> Timeout based with 50 milliseconds of timeout
```
go run *.go pub 
```
- Publisher can request a "List" operation to get the single message state of the entire queue -> list
```
go run *.go sub 
```
- Subscriber can "Check" for instantaneous update also manually -> check




