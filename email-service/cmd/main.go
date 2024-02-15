package main

import (
	"email-service/internal/infrastructure/framework"
	"net/http"
)

func main() {

	app := framework.NewApplication()

	err := http.ListenAndServe(":9977", app)
	if err != nil {
		panic("Error starting email-service application")
	}
}
