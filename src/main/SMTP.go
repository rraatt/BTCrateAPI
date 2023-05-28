package main

import (
	"fmt"
	"net/smtp"
	"os"
	"sync"
)

func sendEmail(to string, rate float64, errChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
		os.Getenv("SMTP_HOST"),
	)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	msg := []byte("To: " + to + "\r\n" +
		"Subject: BTC to UAH rate!\r\n" +
		"\r\n" +
		"At the moment rate is " + fmt.Sprintf("%f", rate) + " UAH for 1 BTC\r\n")
	err := smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_SENDER"), []string{to}, msg)
	if err != nil {
		errChan <- err
		return
	}
}

func sendOutEmails(recipients []string) []error {
	var errs []error
	// Get current rate
	rate, err := fetchRate()
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	// Create channel for catching errors
	errChan := make(chan error, len(recipients))
	var wg sync.WaitGroup
	// Create Goroutines for every email recipient
	for _, recipient := range recipients {
		wg.Add(1)
		go sendEmail(recipient, rate, errChan, &wg)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}
