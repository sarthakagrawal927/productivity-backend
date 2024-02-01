package main

import (
	"todo/cmd/service"
	db "todo/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	service.CreateService()
}

func init() {
	godotenv.Load()
	db.SetupDBConnection()
}
