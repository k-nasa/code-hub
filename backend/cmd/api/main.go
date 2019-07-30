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

	s, err := server.NewServer(datasource)
	if err != nil {
		log.Fatalf("failed to init server: %s", err)
	}
	s.Run(os.Getenv("PORT"))
}
