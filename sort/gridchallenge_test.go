package gridchallenge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestGridCase struct {
	name string
	give []string
	want string
}

func Test2Dto1D(t *testing.T) {

	result := String1Dto2D([]string{"abc", "abc"})
	assert.Equal(t, result, [][]string{[]string{"a", "b", "c"}, []string{"a", "b", "c"}})

}

func TestSortString(t *testing.T) {

	result := SortString([]string{"a", "c", "b"})
	assert.Equal(t, result, []string{"a", "b", "c"})

}

func TestGridChallenge1(t *testing.T) {
	testcase := []struct {
		name string
		give []string
		want string
	}{
		{
			name: "1",
			give: []string{"hcd", "awc", "shm"},
			want: "NO",
		},
		{
			name: "2",
			give: []string{"sur", "eyy", "gxy"},
			want: "NO",
		},
		{
			name: "3",
			give: []string{"nyx", "ynx", "xyt"},
			want: "YES",
		},
		{
			name: "4",
			give: []string{"vpvv", "pvvv", "vzzp", "zzyy"},
			want: "YES",
		},
	}

	for _, tt := range testcase {
		t.Run(tt.name, func(t *testing.T) {
			got := gridChallenge(tt.give)
			if got != tt.want {
				t.Errorf("got %v want %v", got, tt.want)
			}
		})
	}
}
