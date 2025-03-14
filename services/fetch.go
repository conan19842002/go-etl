package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"go-etl/db"
	"go-etl/utils"
)

const API_URL = "https://randomuser.me/api/"

// FetchData retrieves raw user data from the API
func FetchData() []byte {
	resp, err := http.Get(API_URL)
	if err != nil {
		log.Println("Failed to fetch API:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		return nil
	}

	// save raw data to PostgreSQL
	err = db.SaveRawData(string(body))
	if err != nil {
		log.Println("Failed to save raw data:", err)
	}

	// save raw data to JSON file
	utils.SaveToFile("data/raw/raw_data.json", body)

	log.Println("Successfully fetched and stored raw data.")
	return body
}
