package main

import (
	"context"
	"leaderboard/config"
	"leaderboard/controllers"
	"leaderboard/scoredb"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	cfg, err := loadConfig()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.MongoURI).SetServerAPIOptions(serverAPI)
	client := scoredb.Connect(err, opts)
	defer scoredb.Disconnect(err, client)
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	listenAddr := ":" + cfg.ServerPort

	r.GET("/api/info", controllers.Info())

	r.GET("/api/leaderboard", controllers.GetLeaderboard(client, cfg))

	r.GET("/api/webhook", controllers.Webhook())
	r.POST("/api/webhook", controllers.Webhook())

	r.GET("/api/albums", controllers.GetAlbums(client, cfg))
	r.GET("/api/albums/:id", controllers.GetAlbumByID(client, cfg))
	r.POST("/api/albums", controllers.PostAlbum(client, cfg))

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
