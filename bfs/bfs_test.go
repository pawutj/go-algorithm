package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	v := Bfs()
	assert.Equal(t, v[0], 1)
	assert.Equal(t, v[1], 2)
	assert.Equal(t, v[2], 3)
}
