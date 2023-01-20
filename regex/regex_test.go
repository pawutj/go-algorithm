package regex

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex1(t *testing.T) {
	result, err := Regex1()
	assert.Equal(t, result, true)
	assert.Equal(t, err, nil)
}

func TestRegex2(t *testing.T) {
	result, err := Regex2()
	assert.Equal(t, result, true)
	assert.Equal(t, err, nil)
}
