package sqshandlers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitListQueues() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	result, err := GetQueues(sess)
	if err != nil {
		fmt.Println("Got an error retrieving queue URLs:")
		fmt.Println(err)
		return
	}
	for i, url := range result.QueueUrls {
		fmt.Printf("%d: %s\n", i, *url)
	}
}

func GetQueues(sess *session.Session) (*sqs.ListQueuesOutput, error) {
	svc := sqs.New(sess)
	result, err := svc.ListQueues(nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}
