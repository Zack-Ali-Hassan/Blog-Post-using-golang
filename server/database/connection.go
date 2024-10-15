package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDb() *mongo.Database {
	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading environment variables : ", err)
		}
	}

	MONGO_URI := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	fmt.Println("MONGODB connection successfully")
	database := client.Database("golang_blog_db")
	DB = database
	return DB
}
