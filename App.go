package main

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"os"
)

// Constants represents constants provided by the runtime environment
var Constants models.EnvConstants

func main() {
	Constants = models.EnvConstants{
		Token:      os.Getenv("token"),
		BotToken:   os.Getenv("botToken"),
		DBUser:     os.Getenv("dbUser"),
		DBPassword: os.Getenv("dbPassword"),
		DBName:     os.Getenv("dbName"),
		DBHost:     os.Getenv("dbHost"),
	}

	log.Println(Constants)
	err := IntializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	UpdateUserList()
	router := NewRouter()

	fmt.Printf("port: %v", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}
