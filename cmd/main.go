package main

import (
	"CalculatorAPI/application"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // значение по умолчанию
	}
	app := application.New()
	app.RunServer(port)
}
