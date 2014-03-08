package sqsdrain

import (
	"log"
	"time"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
)

const (
	MAX_SQS_MESSAGES = 10
)

func Drain(accessKey string, secretKey string, region aws.Region, queueName string, sleep time.Duration, runForever bool) {
	auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
	client := sqs.New(auth, region)
	queue, err := client.GetQueue(queueName)
	if err != nil {
		log.Fatalf("sqsdrain: could not connect to queue: %s\n", err)
	}
	log.Printf("sqsdrain: connected to queue: %s\n", queue.Url)
	// Drain
	for {
		response, _ := queue.ReceiveMessage(MAX_SQS_MESSAGES)
		if len(response.Messages) == 0 && !runForever {
			break
		}
		for _, m := range response.Messages {
			queue.DeleteMessage(&m)
		}
		term := "message"
		if len(response.Messages) <= 1 {
			term += "s"
		}
		log.Printf("sqsdrain: dequeued %d %s\n", len(response.Messages), term)
		time.Sleep(sleep)
	}
}
