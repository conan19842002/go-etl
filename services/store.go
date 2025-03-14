package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// SaveProcessedData appends new transformed data into a valid JSON array
func SaveProcessedData(newData []byte) error {
	processedFile := "data/processed/processed_data.json"

	// Read existing data (if file exists)
	var existingData []map[string]interface{}
	if _, err := os.Stat(processedFile); err == nil {
		fileContent, err := ioutil.ReadFile(processedFile)
		if err == nil && len(fileContent) > 0 {
			err = json.Unmarshal(fileContent, &existingData)
			if err != nil {
				log.Println("Failed to parse existing JSON data:", err)
				return err
			}
		}
	}

	// Unmarshal new data into an array
	var newRecords []map[string]interface{}
	err := json.Unmarshal(newData, &newRecords)
	if err != nil {
		log.Println("Failed to parse new JSON data:", err)
		return err
	}

	// append new records to the existing data
	existingData = append(existingData, newRecords...)

	// combined data back to JSON
	finalData, err := json.MarshalIndent(existingData, "", "  ")
	if err != nil {
		log.Println("Failed to convert data to JSON:", err)
		return err
	}

	// write updated JSON array back to the file
	err = ioutil.WriteFile(processedFile, finalData, 0644)
	if err != nil {
		log.Println("Failed to save processed data:", err)
		return err
	}

	log.Println("Processed data saved successfully.")
	return nil
}
