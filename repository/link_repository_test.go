package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockRepository struct {

}

func TestNewLinkRepository_CheckRepositoryInstance(t *testing.T) {
	ass := assert.New(t)
	repo := NewLinkRepository()

	ass.IsType(repo , LinkRepository{})
}
