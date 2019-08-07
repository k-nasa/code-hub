package main

import (
	"github.com/joho/godotenv"
	server "github.com/voyagegroup/treasure-app"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	datasource := os.Getenv("DATABASE_DATASOURCE")
	if datasource == "" {
		log.Fatal("Cannot get datasource for database")
	}

	s := server.NewServer()
	s.Init(datasource)
	s.Run(os.Getenv("PORT"))
}
