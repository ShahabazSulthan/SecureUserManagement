package main

import (
	"admin/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" //used for loading environment variables from a .env file.
	"log"
	"fmt"
)

func main() {
	//  To load environment variables from a .env file.
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// initializes a new Gin router with default middleware
	router := gin.Default()

	//passing the router as an argument
	routes.AuthStudentsRoutes(router)
	routes.AuthAdminRoutes(router)

	//starts the Gin router on port 8080
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

