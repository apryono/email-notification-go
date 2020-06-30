package helper

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

/*
	============ EMAIL NOTIFICATION ==============
	We use structure and using api send grid,
	Using feature get http client

	Before using this function, first we initialize environment,
	and we use feature os.Getenv()

	Make variable payload as structure of send grid, and the name is free what you want to initialize,
	we use fmt.Sprintf cause we can make dinamic input on go language



	Note :
	%s => string input
	to => as input for the destination email to be sent
	subject => as a subject in email
	message => as a body on email
	value => as a body input

*/

var (
	authToken        = os.Getenv("AUTH_EMAIL_TOKEN")
	urlEmailSendGrip = os.Getenv("URL_EMAIL_SEND_GRIP")
	methodPost       = "POST" // using method post not get to send the email
	mailSender       = "mail_sender@mail.com"
)

// EmailNotificationSendGrip implement to send email notification
func EmailNotificationSendGrip(to, subject, message string) error {
	payload := fmt.Sprintf(`
		{
			"personalizations":[
				{
					"to": [
						{
							"email": "%s"
						}
					],
					"subject": "%s"
				}
			],
			"from": {
				"email": "%s"
			},
			"content": [
				{
					"type": "text/html",
					"value": "%v"
				}
			]
		}`, to, subject, mailSender, message) // must be sequence

	// and we generate the payload to []byte
	bodyReader := []byte(payload)

	req, err := http.NewRequest(methodPost, urlEmailSendGrip, bytes.NewBuffer(bodyReader))
	if err != nil {
		logrus.Error("Err Request :", err)
		return err // the return we can disable if you wan't to send the error response
	}

	// and then we change and set the header
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer response.Body.Close()
	return nil
}
