package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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

func TestGetURI(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(fmt.Sprintf("mongodb://%s:%s" , os.Getenv("DB_HOST") , os.Getenv("DB_PORT")),GetURI())
}