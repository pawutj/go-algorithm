package regex

import (
	"regexp"
)

func Regex1() (bool, error) {
	s := "seafood"
	matched, err := regexp.MatchString(`foo.*`, s)

	return matched, err

}

func Regex2() (bool, error) {
	s := "pawut.j@gmail.com"
	matched, err := regexp.MatchString(`^.{5,9}\..@.+$`, s)

	return matched, err

}
