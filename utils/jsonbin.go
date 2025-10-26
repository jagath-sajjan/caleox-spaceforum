package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	JSONBIN_URL, API_KEY = GetSecrets() // use embedded secrets
)

// GetBin fetches the latest JSONBin record
func GetBin() (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", JSONBIN_URL+"/latest", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Add("X-Master-Key", API_KEY)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bin: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch bin: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	record, ok := result["record"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid JSONBin structure")
	}

	return record, nil
}

// UpdateBin updates the JSONBin record with new data
func UpdateBin(data map[string]interface{}) error {
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("PUT", JSONBIN_URL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("X-Master-Key", API_KEY)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to update bin: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return fmt.Errorf("failed to update bin: %s", resp.Status)
	}

	return nil
}
