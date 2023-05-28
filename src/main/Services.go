package main

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type PriceData struct {
	Data struct {
		BTC struct {
			Quote struct {
				UAH struct {
					Price float64 `json:"price"`
				} `json:"UAH"`
			} `json:"quote"`
		} `json:"BTC"`
	} `json:"Data"`
}

func getEmails() ([]string, error) {
	filePath := "emails.txt"
	var emails []string
	// Open the file in read-only mode.
	file, err := os.Open(filePath)
	if err != nil {
		return emails, err
	}
	defer file.Close()

	// Create a new scanner and read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		emails = append(emails, strings.TrimSpace(scanner.Text()))
	}
	return emails, err
}

func storeEmail(email string) (bool, error) {
	filePath := "emails.txt"

	// Open the file in read-only mode.
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Create a new scanner and read the file line by line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == email {
			// If the email is found, return false.
			return false, nil
		}
	}

	// If we got here, the email was not found in the file.
	// So, let's append it to the file.
	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	if err != nil {
		return false, err
	}

	return true, nil
}

func fetchRate() (float64, error) {
	apiKey := os.Getenv("API_KEY")
	client := &http.Client{}
	// Create a request
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest", nil)
	if err != nil {
		return 0, err
	}
	// Specify required values
	q := url.Values{}
	q.Add("symbol", "BTC")
	q.Add("convert", "UAH")
	// Authentication for api
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = q.Encode()
	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var PriceData PriceData
	// Deserialize received data to PriceData struct
	err = json.Unmarshal(body, &PriceData)
	if err != nil {
		return 0, err
	}

	return PriceData.Data.BTC.Quote.UAH.Price, err
}
