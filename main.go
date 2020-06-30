package main

import "github.com/email-notification-go/infra/helper"

/*
	====== **** SAMPLE DATA EMAIL NOTIFICATION **** ======

	Depending on how you enter data for the notification email,
	here is just a simple way how to send the notification

	Example for email notifications, when someone wants to apply for a job
	from a career page on a website. And when one of the job positions has been applied for,
	he will receive a notification that the proposed application will be processed.

	========================================================
*/

func main() {

	// we can custom name of the applicant
	candidateName := "John Doe"
	jobPosition := "Project Mananger"
	referenceNumber := "A102"
	candidateEmail := "apryonoboang@gmail.com"

	subject := "Thank you for submitting your job application"
	message := "Dear " + candidateName +
		", <br><br>Thank you for your interest in joining Our Company. Your application for the role of ( " + jobPosition + " ), " +
		"reference number " + referenceNumber + " was submitted successfully.<br><br>" +
		"Just wanted to let you know that we will review your application really soon and we will get back to you with the update.<br><br>" +
		"Have a nice day!"

	err := helper.EmailNotificationSendGrip(candidateEmail, subject, message) // send email notification
	if err != nil {
		return
	}
}
