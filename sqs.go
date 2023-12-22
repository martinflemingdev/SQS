package main

import (
    "github.com/aws/aws-lambda-go/events"
    "time"
)

func main() {
    sqsEvent := events.SQSEvent{
        Records: []events.SQSMessage{
            {
                MessageId:         "example-message-id",
                ReceiptHandle:     "example-receipt-handle",
                Body:              "Hello from SQS!",
                Md5OfBody:         "7b270e59b47ff90a553787216d55d91d",
                Md5OfMessageAttributes: "d3b07384d113edec49eaa6238ad5ff00",
                Attributes: map[string]string{
                    "ApproximateReceiveCount":          "1",
                    "SentTimestamp":                    time.Now().Format(time.RFC3339),
                    "SenderID":                         "123456789012",
                    "ApproximateFirstReceiveTimestamp": time.Now().Format(time.RFC3339),
                },
                MessageAttributes: map[string]events.SQSMessageAttribute{
                    "CustomAttribute": {
                        StringValue:      aws.String("Custom Value"),
                        BinaryValue:      []byte("Binary Value"),
                        StringListValues: []string{"String List Value 1", "String List Value 2"},
                        BinaryListValues: [][]byte{[]byte("Binary List Value 1"), []byte("Binary List Value 2")},
                        DataType:         "String",
                    },
                },
                EventSource:        "aws:sqs",
                EventSourceARN:     "arn:aws:sqs:us-east-1:123456789012:MyQueue",
                AwsRegion:          "us-east-1",
            },
        },
    }

    // Process the SQS event here...
}
