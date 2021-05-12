package main

import (
	"log"
	"os"
	"time"

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

	someData := struct {
		FieldName      string `json:"field_name"`
		IntField       int    `json:"int_field"`
		TimestampField int64  `json:"timestamp_field"`
	}{
		FieldName:      "some_value",
		IntField:       123,
		TimestampField: time.Now().UTC().Unix(),
	}

	// New event
	err = client.NewEventUsingInterface(
		"123", "order_completed", time.Now().UTC(), someData,
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Event Created Successfully!")
}
