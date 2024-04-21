package main

import (
	"fmt"

	"github.com/goropencho/relay/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Environment Not Loaded")
	}
	_ = database.DatabaseConnection()
	fmt.Println("Working Appropriately")
}
