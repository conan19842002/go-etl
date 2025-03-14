package services

import (
	"encoding/json"
	"log"
)

// TransformData processes raw API response and extracts necessary fields
func TransformData(rawData []byte) ([]byte, error) {
	var apiResp map[string]interface{}
	err := json.Unmarshal(rawData, &apiResp)
	if err != nil {
		return nil, err
	}

	results, ok := apiResp["results"].([]interface{})
	if !ok || len(results) == 0 {
		return nil, nil
	}

	transformed := make([]map[string]interface{}, len(results))

	for i, user := range results {
		u := user.(map[string]interface{})
		transformed[i] = map[string]interface{}{
			"name":  u["name"],
			"email": u["email"],
			"dob":   u["dob"].(map[string]interface{})["date"],
		}
	}

	output, err := json.MarshalIndent(transformed, "", "  ")
	if err != nil {
		return nil, err
	}

	log.Println("Data transformed successfully.")
	return output, nil
}
