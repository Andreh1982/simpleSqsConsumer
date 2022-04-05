package worker

import (
	"flag"
	"fmt"
	"time"

	"./worker/sqshandlers"
)

func Start() (queue *string, timeout *int64) {

	queue = flag.String("q", "", "The name of the queue")
	timeout = flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()

	// sqshandlers.InitListQueues()

	for {
		fmt.Println("Waiting For New SQS Messages...")
		time.Sleep(15 * time.Second)
		sqshandlers.InitGetMessages(queue, timeout)
	}
}
