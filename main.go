// Simple HTTP server creation - Entry point for API.
package main

import (
	"log"
	"net/http"
	"os"

	"awesomeProject/controller"
	"github.com/gorilla/handlers"
)

func main()  {
	r := controller.Router()
	log.Fatal(http.ListenAndServe(port(), handlers.CORS()(r)))
}

// Get the port env variable.
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8984"
	}
	return ":" + port
}