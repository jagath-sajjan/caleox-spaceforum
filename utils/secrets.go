package utils

import (
	"encoding/base64"
	"fmt"
)

// 🔒 Embedded Base64-encoded secrets
var (
	encryptedBinURL = "aHR0cHM6Ly9hcGkuanNvbmJpbi5pby92My9iLzY4ZmRiYWFlZDBlYTg4MWY0MGJjOWE0Yg=="
	encryptedAPIKey = "JDJhJDEwJEtQdTEwdHh6MFpTZmhseHhyM2svdHViWk4zaHJETWk5S3ZWRHg1czJLZGtXSWZYM2F3LkR5"
)

// GetSecrets decodes and returns URL and API Key at runtime
func GetSecrets() (string, string) {
	binURLBytes, err := base64.StdEncoding.DecodeString(encryptedBinURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode JSONBin URL: %v", err))
	}

	apiKeyBytes, err := base64.StdEncoding.DecodeString(encryptedAPIKey)
	if err != nil {
		panic(fmt.Sprintf("Failed to decode API Key: %v", err))
	}

	return string(binURLBytes), string(apiKeyBytes)
}
