package crazy

import (
	"errors"
	"strconv"
	"strings"
)

func CrazyString(a string, b string) (string, error) {
	result := ""
	for idx := range a {
		r, err := CrazyRune(rune(a[idx]), rune(b[idx]))
		if err != nil {
			return "", err
		}
		result += string(r)
	}

	return result, nil
}

func CrazyInt(a int64, b int64) (int64, error) {
	a_str := strconv.FormatInt(a, 3)
	b_str := strconv.FormatInt(b, 3)

	if len(a_str) != len(b_str) {
		a_longer := len(a_str) > len(b_str)

		if a_longer {
			l := len(a_str) - len(b_str)
			b_str = strings.Repeat("0", l) + b_str
		} else {
			l := len(b_str) - len(a_str)
			a_str = strings.Repeat("0", l) + a_str
		}
	}

	result_str := ""
	for idx := range a_str {
		r, err := CrazyRune(rune(a_str[idx]), rune(b_str[idx]))
		if err != nil {
			return -1, err
		}
		result_str += string(r)
	}

	return strconv.ParseInt(result_str, 3, 64)
}

var answers = []rune{
	'1', '1', '2', // b == 0
	'0', '0', '2', // b == 1
	'0', '2', '1', // b == 2
}

func CrazyRune(a rune, b rune) (rune, error) {
	if a < '0' || a > '2' {
		return 0x00, errors.New("crazy: a rune not valid")
	}

	if b < '0' || b > '2' {
		return 0x00, errors.New("crazy: b rune not valid")
	}

	index := int(a-'0') + int(b-'0')*3

	return answers[index], nil
}
