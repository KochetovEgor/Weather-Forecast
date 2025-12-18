package main

import (
	"WeatherForecast/app"
	"WeatherForecast/database"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	pool, err := database.NewPool(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		log.Fatal("Error creating databse pool")
	}
	defer pool.Close()
	db := database.NewDB(pool)

	app := app.New(db)
	log.Fatal(app.Run())
}
