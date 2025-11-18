package config

import (
	"context"
	"fmt"
	"jk-api/internal/helper"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	MongoClient *mongo.Client
	MongoDB     *mongo.Database
)

func InitMongoDB() {
	_ = godotenv.Load()

	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DB")

	if uri == "" || dbName == "" {
		helper.LogMessage("ERROR", "❌ MongoDB URI atau DB name tidak ditemukan di environment")
		return
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	var client *mongo.Client
	var err error

	for i := 1; i <= 3; i++ {
		client, err = mongo.Connect(opts)
		if err != nil {
			helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to MongoDB: %v", err))
			time.Sleep(2 * time.Second)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = client.Ping(ctx, readpref.Primary())
		cancel()

		if err != nil {
			log.Printf("⚠️ [MongoDB] Ping percobaan %d gagal: %v", i, err)
			helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to MongoDB: %v", err))
			time.Sleep(2 * time.Second)
			continue
		}

		MongoClient = client
		MongoDB = client.Database(dbName)

		helper.LogMessage("CONFIG", "✅ Successfully connected to MongoDB")
		return
	}

	helper.LogMessage("CONFIG", "✅ Successfully connected to MongoDB")
}

func CloseMongoDB() {
	if MongoClient == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := MongoClient.Disconnect(ctx); err != nil {
		helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to disconnect from MongoDB: %v", err))
	} else {
		helper.LogMessage("CONFIG", "✅ Disconnected from MongoDB")
	}
}
