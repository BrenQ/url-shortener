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
	*mongo.Client
}

// Get a connection database
func GetConnection() *Database {
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	opts := options.Client().ApplyURI(GetURI())

	sess, err := mongo.Connect(ctx, opts)

	if err != nil {
		 log.Fatal("Can´t connect with database" , err)
	}

	return &Database{sess}
}

// Parse URI for mongo connection
func GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s" , os.Getenv("DB_HOST") , os.Getenv("DB_PORT"))
}

/* Approach to considered
func GetInstance () * Database {
	once.Do(func() {
		db = GetConnection()
	})

	return db
}
*/
