package service

import (
	"context"
	"fmt"
	"time"

	"github.com/rickedb/torqueprotracker/concurrent"
	requests "github.com/rickedb/torqueprotracker/v1/api/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var concurrentSlice concurrent.ConcurrentSlice

func Enqueue(entity *requests.TorqueProRequest) {
	concurrentSlice.Append(entity)
}

func Save(database string) {
	documents := make([]interface{}, 0)
	queue := concurrentSlice.PopAll()
	for doc := range queue {
		documents = append(documents, doc)
	}

	if len(documents) > 0 {
		opt := options.Client().ApplyURI(database)
		client, err := mongo.NewClient(opt)
		if err != nil {
			fmt.Printf("Error occurred while creating client: %v", err)
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)
		if err != nil {
			fmt.Printf("Error occurred while connecting: %v", err)
			return
		}

		db := client.Database("torquepro")
		result, err := db.Collection("raw_data").InsertMany(ctx, documents)
		if err != nil {
			fmt.Printf("Error occurred while inserting data: %v", err)
		} else {
			fmt.Printf("Total inserted: %v", len(result.InsertedIDs))
		}
		defer client.Disconnect(ctx)
	}
}
