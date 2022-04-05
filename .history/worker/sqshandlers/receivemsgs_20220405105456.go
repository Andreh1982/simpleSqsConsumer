package sqshandlers

import (
	"alertsplatform-api/core/alert"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitGetMessages(queue *string, timeout *int64) (messageNew alert.AlertSQS, messageRaw string, messageID string, consumedDate string) {

	*queue = "su-sre-alertsapi-tests"
	*timeout = 5

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	urlResult, err := GetQueueURL(sess, queue)

	if urlResult == nil {
		fmt.Println("Not possible to contact remote queue.", err)
		os.Exit(0)
	}

	queueUrl := urlResult.QueueUrl
	msgResult, err := GetMessages(sess, queueUrl, timeout)
	currentDateConsumed := time.Now()

	if err != nil {
		fmt.Println("Got an error receiving messages:")
		fmt.Println(err)
	}

	fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
	fmt.Println("Consumed Date:  " + currentDateConsumed.String())
	json.Unmarshal([]byte(*msgResult.Messages[0].Body), &messageNew)

	messageID = *msgResult.Messages[0].MessageId
	messageRaw = *msgResult.Messages[0].Body
	consumedDate = currentDateConsumed.String()
	receiptHandler := *msgResult.Messages[0].ReceiptHandle

	SendToAlertsApi(messageNew)
	InitDeleteMessage(queueUrl, &receiptHandler, sess)

	return messageNew, messageRaw, messageID, consumedDate
}

func GetQueueURL(sess *session.Session, queue *string) (*sqs.GetQueueUrlOutput, error) {
	svc := sqs.New(sess)
	urlResult, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})
	if err != nil {
		return nil, err
	}
	return urlResult, nil
}

func GetMessages(sess *session.Session, queueURL *string, timeout *int64) (*sqs.ReceiveMessageOutput, error) {
	svc := sqs.New(sess)
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
	})
	if msgResult.Messages == nil {
		fmt.Println("No New Messages Found.")
		os.Exit(0)
	}
	if err != nil {
		return nil, err
	}
	return msgResult, nil
}
