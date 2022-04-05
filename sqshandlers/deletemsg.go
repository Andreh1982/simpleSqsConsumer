package sqshandlers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitDeleteMessage(queueUrl *string, receiptHandler *string, sess *session.Session) {

	err := DeleteMessage(sess, queueUrl, receiptHandler)
	if err != nil {
		fmt.Println("Got an error deleting the message:")
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted message from queue with URL " + *queueUrl)
}

func DeleteMessage(sess *session.Session, queueURL *string, receiptHandle *string) error {
	svc := sqs.New(sess)

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
