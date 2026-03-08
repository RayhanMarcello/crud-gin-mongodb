package main

import (
	"crud-gin-mongodb/db"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	salah := godotenv.Load()
	if salah != nil {
		log.Fatal("gagal load .env")
	}

	_, err := db.NewMongoConfig()
	if err != nil {
		fmt.Println(err)
	}

}
