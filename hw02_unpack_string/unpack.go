package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var ErrInvalidString error

func Unpack(s string) (string, error) {
	var a string
	ra := []rune(s)
	for i, e := range ra {
		if unicode.IsDigit(e) == false {
			a = a + string(e)
		} else if s == "" {
			return "", ErrInvalidString
		} else {
			if i-1 > 0 {
				x, _ := strconv.Atoi(string(e))
				for c := 1; c < x; c++ {
					a = a + string(ra[i-1])
				}
			} else if unicode.IsDigit(ra[i-1]) == true {
				ErrInvalidString = errors.New("invalid string")
			} else {
				ErrInvalidString = errors.New("invalid string")
			}
		}
	}
	fmt.Print(a)
	return a, ErrInvalidString
}
