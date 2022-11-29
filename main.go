package main

import (
	db "www.github.com/N4choWasTaken/Go-API-JWT/database"
)

func main() {
	// Initialize Database
	db.Connect("host=localhost user=postgres password=mysecretpassword dbname=postgres port=5455 sslmode=disable")
	db.Migrate()
}
