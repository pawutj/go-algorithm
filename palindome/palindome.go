package palindome

func delChar(s []rune, index int) string {
	return string(append(s[0:index], s[index+1:]...))
}

func CheckPalindome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func PalindomeRemove(s string) int {

	if CheckPalindome(s) {
		return -1
	}

	for i := 0; i < len(s); i++ {
		s2 := []rune(s)
		res := delChar(s2, i)
		if CheckPalindome(res) {
			return i
		}
	}

	return -1

}
