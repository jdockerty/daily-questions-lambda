package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func handleRequest() (string, error) {
	newSession := session.Must(session.NewSession())
	sesClient := ses.New(newSession, aws.NewConfig().WithRegion("us-west-2"))
	emailContent := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String("MY_EMAIL")},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("HTML question format here, same as Python version."),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("Message plain text here"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Daily Questions"),
			},
		},
		Source: aws.String("MY_EMAIL"),
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
