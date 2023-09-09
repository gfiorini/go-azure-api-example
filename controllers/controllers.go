package controllers

import (
	"io"
	"leaderboard/config"
	"leaderboard/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Webhook() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		sb := string(body)
		log.Println(sb)
		log.Println(c.Request.URL.Query())
	}
}

func Info() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo: recuperare da file di build
		var info = model.Info{Version: "0.0.1"}
		c.IndentedJSON(http.StatusOK, info)
	}
}

func GetScores(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		var scores []model.Score
		cursor, _ := coll.Find(ctx, bson.M{})
		cursor.All(ctx, &scores)
		var leaderboard model.ScoresContainer
		leaderboard.Scores = scores
		ctx.IndentedJSON(http.StatusOK, leaderboard)
	}
}

func PostScore(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var score model.Score
		err := ctx.BindJSON(&score)
		if err != nil {
			log.Printf("Error: %v", err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
		}
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		if res, err := coll.InsertOne(ctx, &score); err != nil {
			log.Printf("Error: %v", err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
		} else {
			log.Printf("Inserted record with id: %v", res.InsertedID)
			ctx.IndentedJSON(http.StatusCreated, score)
		}
	}
}

func GetScoreByID(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		var score model.Score
		err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&score)
		if err != nil {
			log.Printf("Error: %v", err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
		}
		ctx.IndentedJSON(http.StatusOK, score)
	}
}

func DeleteAllScores(mongoClient *mongo.Client, cfg config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbScoresCollection)
		if _, err := coll.DeleteMany(ctx, bson.D{}); err != nil {
			log.Println(err)
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "generic error"})
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{})
		}
	}
}
