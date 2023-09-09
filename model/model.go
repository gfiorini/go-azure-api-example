package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type Info struct {
	Version string `json:"version"`
}

type ScoresContainer struct {
	Scores []Score `json:"scores" bson:"scores"`
}
type Score struct {
	Name     string             `json:"name" bson:"name"`
	Value    int                `json:"value" bson:"value"`
	PlayerID string             `json:"playerID" bson:"playerID"`
	Date     primitive.DateTime `json:"date" bson:"date"`
}
