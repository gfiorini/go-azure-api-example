package scoredb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leaderboard/config"
	"log"
)

func Disconnect(err error, client *mongo.Client) {
	func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
}

func Connect(err error, opts *options.ClientOptions) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal("cannot connect to mongo:", err)
	}
	return client
}

func Ping(client *mongo.Client, config config.Config) {
	if err := client.Database(config.MongoDbScoreDatabase).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
