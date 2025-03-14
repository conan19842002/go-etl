package main

import (
	"go-etl/config"
	"go-etl/db"
	"go-etl/server"
	"go-etl/services"
	"go-etl/utils"
	"log"
	"time"
)

func main() {
	// load configuration
	config.LoadConfig()

	// connect to PostgreSQL
	db.ConnectDB()

	// initialize logger
	utils.InitLogger()

	// start the HTTP server
	log.Println("Starting HTTP server on port 8080...")
	go func() {
		server.StartServer()
	}()

	// Log message to confirm main.go is running
	log.Println("Main function is running!")

	// Create a channel for passing raw data
	dataChannel := make(chan []byte, 5)

	// Start data processing worker
	go processWorker(dataChannel)

	// Fetch data every 30 seconds and push to the channel
	for {
		rawData := services.FetchData()
		if rawData == nil {
			log.Println("Failed to fetch data. Skipping...")
		} else {
			dataChannel <- rawData
		}
		time.Sleep(30 * time.Second)
	}
}

// processWorker continuously transforms and saves incoming data
func processWorker(dataChannel <-chan []byte) {
	for rawData := range dataChannel {
		transformedData, err := services.TransformData(rawData)
		if err != nil {
			log.Println("Transformation failed:", err)
			continue
		}

		err = services.SaveProcessedData(transformedData)
		if err != nil {
			log.Println("Failed to save processed data:", err)
			continue
		}

		log.Println("Successfully processed and saved transformed data.")
	}
}
