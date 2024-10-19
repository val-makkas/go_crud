package main

import (
	"crud/crud_api/db"
	"crud/crud_api/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
