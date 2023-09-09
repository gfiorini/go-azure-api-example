package controllers

import (
	"fmt"
	"io"
	"leaderboard/config"
	"leaderboard/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PostAlbum(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var album model.Album
		var err error
		if album, err = postSingleAlbum(ctx, mongoClient, cfg); err != nil {
			fmt.Println(err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
			return
		}
		ctx.IndentedJSON(http.StatusCreated, album)
	}
}

func postSingleAlbum(ctx *gin.Context, mongoClient *mongo.Client, cfg config.Config) (model.Album, error) {
	var album model.Album
	coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbAlbumsCollection)
	if err := ctx.BindJSON(&album); err != nil {
		return album, err
	}
	if _, err := coll.InsertOne(ctx, &album); err != nil {
		return album, err
	}
	return album, nil
}

func GetAlbums(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := getAllAlbums(ctx, mongoClient, cfg)
		if err != nil {
			fmt.Println(err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "generic error"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func getAllAlbums(ctx *gin.Context, mongoClient *mongo.Client, cfg config.Config) ([]bson.M, error) {
	coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbAlbumsCollection)
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var albums []bson.M
	if err = cursor.All(ctx, &albums); err != nil {
		log.Fatal(err)
	}
	return albums, err
}

func GetAlbumByID(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		res, err := getAlbum(ctx, mongoClient, cfg, id)
		if err != nil {
			//todo: usare logger
			fmt.Println("error occured while fetching album")
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, res)
	}
}

func getAlbum(ctx *gin.Context, mongoClient *mongo.Client, cfg config.Config, id string) (model.Album, error) {
	coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbAlbumsCollection)
	var album model.Album
	err := coll.FindOne(ctx, bson.D{{"id", id}}).Decode(&album)
	return album, err
}

func Webhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		sb := string(body)
		fmt.Println(sb)
		fmt.Println(c.Request.URL.Query())
	}
}

func Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo: recuperare da file di build
		var info = model.Info{Version: "0.0.1"}
		c.IndentedJSON(http.StatusOK, info)
	}
}

func GetLeaderboard(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		var scores []model.Score
		cursor, _ := coll.Find(ctx, bson.M{})
		cursor.All(ctx, &scores)
		var leaderboard model.Leaderboard
		leaderboard.Scores = scores
		ctx.IndentedJSON(http.StatusOK, leaderboard)
	}
}

func PostScore(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var score model.Score
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		if res, err := coll.InsertOne(ctx, &score); err != nil {
			fmt.Printf("Error: %v", err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
		} else {
			fmt.Printf("Inserted record with id: %v", res.InsertedID)
			ctx.IndentedJSON(http.StatusCreated, score)
		}
	}
}
