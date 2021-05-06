package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with Tracking API & App API enabled)
	client, err := customerio.NewClient(
		customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
		customerio.WithAppKey(os.Getenv("APP_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Start the email request
	emailRequest := &customerio.EmailRequest{
		Body:        "This is an example body!",
		From:        "noreply@example.com",
		Identifiers: map[string]string{"id": "123"},
		MessageData: map[string]interface{}{
			"name": "Person",
			"items": map[string]interface{}{
				"name":  "shoes",
				"price": "59.99",
			},
			"products": []interface{}{},
		},
		PlaintextBody: "This is an example body!",
		ReplyTo:       "noreply@example.com",
		Subject:       "Customer io test email",
		To:            "bob@example.com",
	}

	// Attach a file (example)
	/*
		var f *os.File
		f, err = os.Open("<path to file>")
		if err != nil {
			log.Fatalln(err)
		}
		defer func() {
			_ = f.Close()
		}()
		if err = emailRequest.Attach("sample.pdf", f); err != nil {
			log.Fatalln(err)
		}
	*/

	// Send an email (NOT using a template)
	if _, err = client.SendEmail(emailRequest); err != nil {
		log.Fatalln(err)
	}
	log.Println("Email Sent Successfully!")
}
