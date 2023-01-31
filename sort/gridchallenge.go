package gridchallenge

import (
	"sort"
)

func String1Dto2D(grid []string) [][]string {
	string2D := [][]string{}
	for i := 0; i < len(grid); i++ {
		t := []string{}
		for j := 0; j < len(grid[0]); j++ {
			t = append(t, string(grid[i][j]))
		}
		string2D = append(string2D, t)
	}
	return string2D
}

func SortString(a []string) []string {
	temp := append([]string{}, a...)
	sort.Strings(temp)
	return temp
}
func gridChallenge(grid []string) string {

	string2D := String1Dto2D(grid)

	for i := 0; i < len(string2D); i++ {
		string2D[i] = SortString(string2D[i])
	}

	for i := 0; i < len(string2D[0]); i++ {
		temp := []string{""}
		for j := 0; j < len(string2D); j++ {
			temp = append(temp, string2D[j][i])
		}
		sorted := SortString(temp)

		for j := 0; j < len(sorted); j++ {
			if temp[j] != sorted[j] {
				return "NO"
			}
		}
	}

	return "YES"

}
