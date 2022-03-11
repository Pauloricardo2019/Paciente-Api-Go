package main

import (
	"GoPacientes/src/database"
	"GoPacientes/src/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
