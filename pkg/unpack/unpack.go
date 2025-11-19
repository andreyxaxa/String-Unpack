package unpack

import (
	"errors"
	"unicode"
)

// String .
func String(s string) (string, error) {
	// пустая
	if len(s) == 0 {
		return "", nil
	}

	// первый символ - цифра
	if s[0] >= '0' && s[0] <= '9' {
		return "", errors.New("invalid string")
	}

	var res []rune
	var prev rune
	hasPrev := false
	escaped := false

	for _, r := range s {
		if escaped {
			res = append(res, r)
			prev = r
			hasPrev = true
			escaped = false
			continue
		}

		if r == '\\' {
			escaped = true
			continue
		}

		// если цифра
		if unicode.IsDigit(r) {
			if !hasPrev {
				return "", errors.New("digit without pre-char")
			}
			// если перед цифрой была буква - нужно клонировать
			count := int(r - '0')
			for i := 1; i < count; i++ {
				res = append(res, prev)
			}
			continue
		}

		// обычный символ
		res = append(res, r)
		prev = r
		hasPrev = true
	}

	if escaped {
		return "", errors.New("\\ at the end")
	}

	return string(res), nil
}
