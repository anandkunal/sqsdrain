# sqsdrain

sqsdrain is a Go package that removes all messages from an SQS queue. It's incredibly useful for test environments where purging a queue is necessary.


## Installation

Run the following command from your terminal: 

	go get http://github.com/anandkunal/sqsdrain

This package requires goamz, particularly the repository managed by CrowdMob. You can get that dependency by executing the following terminal commands:

	go get github.com/crowdmob/goamz/aws
	go get github.com/crowdmob/goamz/sqs


## Example

Go ahead and couch the following into an `example.go`.

	package main
	
	import (
	  "time"
	
	  "github.com/anandkunal/sqsdrain"
	  "github.com/crowdmob/goamz/aws"
	)
	
	var (
	  accessKey  = "access key goes here"
	  secretKey  = "secret key goes here"
	  region     = aws.USWest2
	  queueName  = "queue-name-goes-here"
	  sleep      = time.Second * 5
	  runForever = true
	)
	
	func main() {
	  sqsdrain.Drain(accessKey, secretKey, region, queueName, sleep, runForever)
	}

You can run the above with a simple `go run example.go`. You should see output like:

	2014/03/07 21:37:59 sqsdrain: connected to queue: https://sqs.us-west-2.amazonaws.com/48239042/queue-name-goes-here
	2014/03/07 21:37:59 sqsdrain: dequeued 0 messages
	2014/03/07 21:38:05 sqsdrain: dequeued 4 message
	2014/03/07 21:38:10 sqsdrain: dequeued 0 messages


## Contributions

sqsdrain was developed by [Kunal Anand][0] and is completely free. Have fun.


[0]: https://twitter.com/ka
[2]: http://mit-license.org/