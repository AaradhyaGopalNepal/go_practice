package mongodb

import (
	"context"
	"gRPC_API/pkg/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoClient(ctx context.Context) (*mongo.Client, error) {
	//"username:password@mongodb://localhost:27017"
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to connect to database")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, utils.ErrorHandler(err, "Unable to ping")

	}
	return client, nil
}
