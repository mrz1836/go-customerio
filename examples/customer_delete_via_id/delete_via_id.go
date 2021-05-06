package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with Tracking API enabled)
	client, err := customerio.NewClient(
		customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Deleting the customer
	err = client.DeleteCustomer("123")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Customer Deleted Successfully!")
}
