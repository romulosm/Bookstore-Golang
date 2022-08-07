package main

import (
	"livraria-golang/database"
	"livraria-golang/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()

}
