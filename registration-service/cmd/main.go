package main

import (
	"net/http"
	"registration-service/internal/infrastructure/framework"
)

func main() {
	app := framework.NewApplication()

	err := http.ListenAndServe(":9099", app)
	if err != nil {
		panic("Error starting registration-service application")
	}

}
