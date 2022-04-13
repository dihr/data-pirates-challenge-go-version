package main

import (
	"data-pirates-challenge-go-version/services/correiosapi"
	"data-pirates-challenge-go-version/services/datawriter"
	"data-pirates-challenge-go-version/services/scraper"
	"log"
)

func main() {
	// Starts api services.
	log.Println("Starting the api services...")
	apiSvc := correiosapi.NewCorreiosApiService()

	// Starts data writer services.
	log.Println("Starting the writer services...")
	writerSvc := datawriter.NewDataWriter()

	// Starts scraper services
	log.Println("Starting the scraper services...")
	scraperSvc := scraper.NewScraperSvc(apiSvc, writerSvc)

	// Runs scraper
	log.Println("Data scraping...")
	if err := scraperSvc.StartScraping(); err != nil {
		panic(err)
	}
}
