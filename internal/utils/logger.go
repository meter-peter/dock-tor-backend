package utils

import (
	"log"
)

func LogError(message string) {
	log.Printf("[ERROR] %s", message)
}

func LogInfo(message string) {
	log.Printf("[INFO] %s", message)
}

func LogWarning(message string) {
	log.Printf("[WARNING] %s", message)
}
