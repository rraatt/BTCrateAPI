package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	e.GET("/rate", getRate)
	e.POST("/subscribe", addEmail)
	e.POST("/sendEmails", sendRates)

	e.Start(":8080")
}
