package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/RHGHospitality/db"
	"github.com/raghavendrah25/RHGHospitality/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://administrator:mongopass@localhost:27017" // Replace with your MongoDB URI

func main() {

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	defer client.Disconnect(ctx)

	h := handlers.NewUserHandler(db.NewMongoDBUserStore(client))

	listenAddress := flag.String("port", ":8080", "Port to run the server on")
	flag.Parse()

	router := gin.Default()
	router.GET("/health", handlers.HandleHealthCheck)

	apiV1 := router.Group("/api/v1")
	apiV1.GET("/about", handlers.HandleAbout)
	apiV1.GET("/user/:id", h.GetUserByID)

	err = router.Run(*listenAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
