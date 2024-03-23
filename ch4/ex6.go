package ch4

import "unicode"

func SquashSpace(slice []byte) []byte {
	n := len(slice)

	for i := 0; i < n; i++ {
		ch := slice[i]

		if unicode.IsSpace(rune(ch)) {
			j := i + 1
			for ; j < n && unicode.IsSpace(rune(slice[j])); j++ {

			}
			slice[i] = ' '
			copy(slice[i+1:], slice[j:])
			n -= j - i - 1
			slice = slice[0:n]
		}
	}

	return slice
}
