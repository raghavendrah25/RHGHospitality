package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/RHGHospitality/handlers"
)

func main() {
	listenAddress := flag.String("port", ":8080", "Port to run the server on")
	flag.Parse()

	router := gin.Default()
	router.GET("/health", handlers.HandleHealthCheck)

	apiV1 := router.Group("/api/v1")
	apiV1.GET("/about", handlers.HandleAbout)
	apiV1.GET("/user", handlers.HandleGetUser)
	apiV1.GET("/user/:id", handlers.HandleGetUserByID)

	err := router.Run(*listenAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
