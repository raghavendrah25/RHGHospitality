package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/raghavendrah25/RHGHospitality/handlers"
	"github.com/raghavendrah25/RHGHospitality/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://administrator:mongopass@localhost:27017" // Replace with your MongoDB URI
const dbname = "rhghospitality-dev"                               // Replace with your database name
const userCollection = "users"                                    // Replace with your collection name

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
	ctx := context.Background()
	coll := client.Database(dbname).Collection(userCollection)

	user := types.User{
		ID:        primitive.NewObjectID().Hex(),
		FirstName: "John",
		LastName:  "Doe",
	}
	result, err := coll.InsertOne(ctx, user)
	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}
	fmt.Println("Inserted document with ID:", result.InsertedID)

	var fetchedUser types.User
	res := coll.FindOne(ctx, map[string]interface{}{"_id": user.ID}).Decode(&fetchedUser)
	if res != nil {
		fmt.Println("Error fetching document:", res)
		return
	}
	fmt.Printf("Fetched User: %+v\n", fetchedUser)

	listenAddress := flag.String("port", ":8080", "Port to run the server on")
	flag.Parse()

	router := gin.Default()
	router.GET("/health", handlers.HandleHealthCheck)

	apiV1 := router.Group("/api/v1")
	apiV1.GET("/about", handlers.HandleAbout)
	apiV1.GET("/user", handlers.HandleGetUser)
	apiV1.GET("/user/:id", handlers.HandleGetUserByID)

	err = router.Run(*listenAddress)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
