package services

import (
	"encoding/json"
	"go-etl/utils"
)

func TransformData(input []byte) ([]byte, error) {
	var apiResp map[string]interface{}
	err := json.Unmarshal(input, &apiResp)
	if err != nil {
		return nil, err
	}

	// Extract needed fields
	results := apiResp["results"].([]interface{})
	transformed := make([]map[string]interface{}, len(results))

	for i, user := range results {
		u := user.(map[string]interface{})
		transformed[i] = map[string]interface{}{
			"name":      u["name"],
			"email":     u["email"],
			"dob":       u["dob"].(map[string]interface{})["date"],
		}
	}

	output, _ := json.MarshalIndent(transformed, "", "  ")
	utils.SaveToFile("data/processed/processed_data.json", output)
	return output, nil
}
