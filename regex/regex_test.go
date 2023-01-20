package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	result, err := Regex()
	assert.Equal(t, result, true)
	assert.Equal(t, err, nil)
}
