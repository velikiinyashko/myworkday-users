package mongodb

import (
	"context"
	"fmt"
	"log"

	"github.com/velikiinyashko/myworkday/pkg/client"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoClient(ctx context.Context, data client.ConnectionData) (db *mongo.Database, err error) {
	connectionStr := fmt.Sprintf("mongodb://%s:%s", data.Server, data.Port)
	clientOpt := options.Client().ApplyURI(connectionStr)

	if data.User != "" && data.Password != "" {
		credential := options.Credential{
			Username: data.User,
			Password: data.Password,
		}
		clientOpt.SetAuth(credential)
	}
	client, err := mongo.Connect(ctx, clientOpt)

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(data.Base), nil

}
