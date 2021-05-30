package main

import (
	"log"
	"os"
	"time"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with Tracking API & App API enabled)
	client, err := customerio.NewClient(
		customerio.WithAppKey(os.Getenv("APP_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.UpdateCollection(
		"",
		"test_collection",
		[]map[string]interface{}{
			{
				"item_name":       "test_item_1",
				"id_field":        1,
				"timestamp_field": time.Now().UTC().Unix(),
			},
			{
				"item_name":       "test_item_2",
				"id_field":        2,
				"timestamp_field": time.Now().UTC().Unix(),
			},
		})
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("Collection Added Successfully!")
}
