package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDb struct {
	coll *mongo.Collection
}

var (
	mongodb *MongoDb
	upsert bool
)

func (mongodb *MongoDb) Initialize() {
	c, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	clientOptions := options.ClientOptions{}
	clt, err := mongo.Connect(c, clientOptions.ApplyURI(config.uri))
	if err != nil {
		checkErr(err, "cannot connect to mongodb")
	}
	logger.Info("successfully connect to mongodb")
	mongodb.coll = clt.Database(config.db).Collection(config.coll)
}