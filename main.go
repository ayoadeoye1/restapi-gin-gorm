package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ayoadeoye1/restapi-gin-gorm/helper"
	"github.com/ayoadeoye1/restapi-gin-gorm/router"
	"github.com/joho/godotenv"

	//docs
	_ "github.com/ayoadeoye1/restapi-gin-gorm/docs"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	routes := router.SetupRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: routes,
	}

	fmt.Printf("\n Server listening on: http://localhost:%d \n", port)

	err = server.ListenAndServe()
	helper.ErrorPanic(err)
}
