package services

import (
    "encoding/json"
    "os"
)

// SaveToJSON writes structured data to a JSON file
func SaveToJSON(filename string, data interface{}) error {
    file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    jsonData, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }

    _, err = file.Write(append(jsonData, '\n'))
    return err
}
