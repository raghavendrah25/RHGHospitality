package db

import (
	"context"

	"github.com/raghavendrah25/RHGHospitality/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const dbname = "rhghospitality-dev" // Replace with your database name
const userCollection = "users"      // Replace with your collection name

type UserStore interface {
	GetUserByID(ctx context.Context, id string) (*types.User, error)
}	

type MongoDBUserStore struct {
	coll   *mongo.Collection
}

func NewMongoDBUserStore(client *mongo.Client) *MongoDBUserStore {
	return &MongoDBUserStore{
		coll:   client.Database(dbname).Collection(userCollection),
	}
}

func (store *MongoDBUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	err := store.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
