package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	data "github.com/thegeniusgroup/SonuKanwar/Data"
)

func main() {
	// Check DynamoDB
	_, err := data.InitDynamoDB()
	if err != nil {
		log.Panic(err)
	}

	// Set up API routes
	router := httprouter.New()
	router.RedirectTrailingSlash = true
	addRouteHandlers(router)

	fmt.Println("Setup complete. Running API server...")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "Authorization"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
