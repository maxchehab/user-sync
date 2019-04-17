package main

import (
	"log"
	"models"
	"net/http"
	"os"
)

// Constants represents constants provided by the runtime environment
var Constants models.EnvConstants

func main() {
	Constants = models.EnvConstants{
		Token:    os.Getenv("token"),
		BotToken: os.Getenv("botToken"),
		DBUrl:    os.Getenv("DATABASE_URL"),
		Port:     os.Getenv("PORT"),
		APIKey:   os.Getenv("apiKey"),
	}

	err := IntializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")

	UpdateUserList()
	log.Println("User list updated")

	router := NewRouter()
	log.Println("Listening >")
	log.Fatal(http.ListenAndServe(":"+Constants.Port, router))
}
