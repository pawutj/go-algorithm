package palindome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindomeCheck1(t *testing.T) {
	result := CheckPalindome("abcba")
	assert.Equal(t, result, true)
}

func TestPalindomeCheck2(t *testing.T) {
	result := CheckPalindome("abcb")
	assert.Equal(t, result, false)
}

func TestPalindomeRemove(t *testing.T) {
	result := PalindomeRemove("abcdba")
	assert.Equal(t, result, 3)

}
