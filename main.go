package main

import (
	"strateegy/project-service/database"
	"strateegy/project-service/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}
