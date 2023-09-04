package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"leaderboard/config"
	"leaderboard/model"
	"log"
	"net/http"
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

func getAlbum(ctx *gin.Context, mongoClient *mongo.Client, cfg config.Config, id string) (bson.D, error) {
	coll := mongoClient.Database(cfg.MongoDbScoreDatabase).Collection(cfg.MongoDbAlbumsCollection)
	var result bson.D
	err := coll.FindOne(ctx, bson.D{{"id", id}}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}
