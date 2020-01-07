package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest() (string, error) {
	newSession := session.Must(session.NewSession())
	sesClient := ses.New(newSession, aws.NewConfig().WithRegion("us-west-2"))
	emailContent := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String("jdockerty19@gmail.com")},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("'text here'"),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This is the message body in text format."),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Daily Questions"),
			},
		},
		Source:        aws.String("jdockerty19@gmail.com"),
	}

	result, err := sesClient.SendEmail(emailContent)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return "Lambda executed", nil
}
func main() {
	lambda.Start(handleRequest)
	
}
