package main

import (
	"HR_management_system/database"
	"HR_management_system/routes"
)

func main() {
	r := routes.SetupRouter()
	dataSourceName := "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
	database.InitDB(dataSourceName)
	r.Run(":8080")
}
