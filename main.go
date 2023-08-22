package main

import (
	"crud-api-barang/config"
	"crud-api-barang/models"
	"crud-api-barang/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := config.InitDB()

	err := db.AutoMigrate(&models.Barang{})
	if err != nil {
		fmt.Println("Error creating table:", err)
	} else {
		fmt.Println("Table barang created successfully")
	}

	r := routes.SetupRouter(db)

	fmt.Println("server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
