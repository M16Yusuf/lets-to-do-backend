package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"

	"github.com/m16yusuf/lets-to-do/internal/configs"
	"github.com/m16yusuf/lets-to-do/internal/routers"
)

func main() {
	// Inisialization databae for this project
	db, err := configs.InitDB()
	if err != nil {
		log.Println("FAILED CONFIG DB")
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed connect to DB: %v", err)
	} else {
		log.Println("Database connect successfully")
	}
	defer sqlDB.Close()

	// initizalie router
	router := routers.InitRouter(db)
	router.Run("127.0.0.1:3000")
}
