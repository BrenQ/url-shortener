package database

import "sync"

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	//"sync"
	"time"
)

// Database connection
var db *Database

// Another approach
var once sync.Once

// Store a mongo client

type Database struct {
	*mongo.Collection
}

// Get a connection database
func NewConnection() *Database {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	opts := options.Client().ApplyURI(GetURI())

	sess, err := mongo.Connect(ctx, opts)

	if err != nil {
		 log.Fatal("Can´t connect with database" , err)
	}

	db := sess.Database("url").Collection("link")

	return &Database{db}
}

// Parse URI for mongo connection
func GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s" , os.Getenv("DB_HOST") , os.Getenv("DB_PORT"))
}

// Approach to be considered.
// I need to check if the following approach like a "singleton" is a good choice.

// func GetInstance () * Database {
//	once.Do(func() {
//		db = GetConnection()
//	})
//
//	return db
//}

