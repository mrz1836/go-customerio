package main

import (
	"log"
	"os"

	"github.com/mrz1836/go-customerio"
)

func main() {

	// Load the client (with Tracking)
	client, err := customerio.NewClient(
		customerio.WithTrackingKey(os.Getenv("TRACKING_SITE_ID"), os.Getenv("TRACKING_API_KEY")),
	)
	if err != nil {
		log.Fatalln(err)
	}

	// Find the region
	var region *customerio.RegionInfo
	region, err = client.FindRegion()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Region found! %+v\\n", region)
}
