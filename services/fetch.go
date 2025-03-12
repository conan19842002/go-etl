package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"go-etl/db"
	"go-etl/utils"
)

const API_URL = "https://randomuser.me/api/"

type APIResponse struct {
	Results []map[string]interface{} `json:"results"`
}

func FetchData() {
	for {
		resp, err := http.Get(API_URL)
		if err != nil {
			log.Println("Failed to fetch API:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Println("Failed to read response body:", err)
			continue
		}

		// Save raw data to PostgreSQL
		err = db.SaveRawData(string(body))
		if err != nil {
			log.Println("Failed to save raw data:", err)
		}

		// Save raw data to JSON file
		utils.SaveToFile("data/raw/raw_data.json", body)

		log.Println("Successfully fetched and stored data.")
		time.Sleep(30 * time.Second)
	}
}
