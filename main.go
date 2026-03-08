package main

import (
	"crud-gin-mongodb/db"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	salah := godotenv.Load()
	if salah != nil {
		log.Fatal("gagal load .env")
	}

	uri := os.Getenv("MONGO_URI")
	fmt.Println(uri)

	_, err := db.NewMongoConfig()
	if err != nil {
		fmt.Println(err)
	}

}
