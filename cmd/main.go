package main

import (
	"database/sql"
	"log"

	database "github.com/Ilhamkawe/grpc-campaign-service/internal/adapter/database/postgres"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://postgres:kawe123@localhost:5433/crowdfunding?sslmode=disable")

	if err != nil {
		log.Fatalln("Cant connect database : ", err)
	}

	databaseAdapter, err := database.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Cant create database adapter : ", err)
	}

}
