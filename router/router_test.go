package router

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRouterNew_CheckRouterInstance(t *testing.T) {
	ass := assert.New(t)
	router := New()

	ass.IsType(router , Router{})
}
