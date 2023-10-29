package main

import (
	"todo/cmd/service"
	db "todo/pkg/database"
)

func main() {
	service.CreateService()
}

func init() {
	db.SetupDBConnection()
}
