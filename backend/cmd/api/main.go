package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	server "github.com/voyagegroup/treasure-app"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file. %s", err)
	}

	datasource := os.Getenv("DATABASE_DATASOURCE")
	if datasource == "" {
		log.Fatal("Cannot get datasource for database.")
	}

	s := server.NewServer()
	s.Init(datasource)
	s.Run(os.Getenv("PORT"))
}
