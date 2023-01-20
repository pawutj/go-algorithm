package regex

import (
	"regexp"
)

func Regex() (bool, error) {
	matched, err := regexp.Match(`foo.*`, []byte(`seafood`))

	return matched, err

}
