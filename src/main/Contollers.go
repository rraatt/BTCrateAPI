package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func getRate(c echo.Context) error {
	price, err := fetchRate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]float64{"status": price})
}

func addEmail(c echo.Context) error {
	email := c.FormValue("email")
	success, err := storeEmail(email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if success {
		return c.JSON(http.StatusOK, map[string]string{"status": "success"})
	} else {
		return c.JSON(http.StatusConflict, map[string]string{"status": "error"})
	}
}

func sendRates(c echo.Context) error {
	emails, err := getEmails()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	errs := sendOutEmails(emails)
	if len(errs) > 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while sending emails"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}
