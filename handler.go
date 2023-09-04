package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"leaderboard/config"
	"leaderboard/model/albums"
	"leaderboard/scoredb"
	"log"
)

var client *mongo.Client

func main() {
	r := gin.Default()
	cfg, err := loadConfig()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoDbConnection).SetServerAPIOptions(serverAPI)
	client := scoredb.Connect(err, opts)
	defer scoredb.Disconnect(err, client)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := ":" + cfg.ServerPort
	r.GET("/api/albums", albums.GetAlbums)
	r.GET("/api/albums/:id", albums.GetAlbumByID)
	r.POST("/api/albums", albums.PostAlbums)

	//r.GET("/api/scores", GetScores)
	r.Run(listenAddr)

}

func loadConfig() (config.Config, error) {
	cfg, err := config.LoadConfig("./config/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return cfg, err
}
