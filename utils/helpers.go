package utils

import (
	"fmt"
	"log"
	"os"
)

// SaveToFile writes data to a file (append mode)
func SaveToFile(filename string, data []byte) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	_, err = f.Write(append(data, '\n'))
	if err != nil {
		log.Println("Error writing to file:", err)
	}
}

// Logger
func InitLogger() {
	file, err := os.OpenFile("logs/etl.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		fmt.Println("Failed to log to file, using default stderr")
	}
}
