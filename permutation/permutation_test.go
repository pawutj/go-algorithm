package permutations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermu(t *testing.T) {
	p := PermutationObject{Count: 0}
	p.permu([]int{})

	assert.Equal(t, p.Count, 64)
}
