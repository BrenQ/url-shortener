package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	c "urlshortener/config"
)

func TestNewConnection_CheckDiffInstance(t *testing.T) {
	ass := assert.New(t)
	// Currently are different instances
	conn1 := NewConnection()
	conn2 := NewConnection()
	ass.NotEqual(conn1,conn2)
}

func TestNewConnection_CheckDatabaseInstance(t *testing.T) {
	ass := assert.New(t)
	// Currently are different instances
	conn := NewConnection()
	ass.IsType(&Database{},conn)
}

func TestLocalGetURI(t *testing.T) {
	ass := assert.New(t)

	cfg := c.Config{
		Env:    "local",
		Prefix: "../.env.",
	}

	cfg.Start()
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s" ,
		cfg.Get("DB_USER"),
		cfg.Get("DB_PASSWORD"),
		cfg.Get("DB_HOST") ,
		cfg.Get("DB_PORT"))

	ass.Equal(uri,GetURI())
}