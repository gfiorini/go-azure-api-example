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

	r.GET("/api/scores/:id", controllers.GetScoreByID(client, cfg))
	r.GET("/api/scores", controllers.GetScores(client, cfg))
	r.POST("/api/scores", controllers.PostScore(client, cfg))
	r.DELETE("/api/scores", controllers.DeleteAllScores(client, cfg)) //non disponibile nelle functions

	r.GET("/api/webhook", controllers.Webhook())
	r.POST("/api/webhook", controllers.Webhook())

	r.Run(listenAddr)

}

func loadConfig() (config.Config, error) {
	cfg, err := config.LoadConfig("./config/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	return cfg, err
}
