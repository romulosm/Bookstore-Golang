package main

import (
	"bookstore-golang/database"
	"bookstore-golang/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()

}
