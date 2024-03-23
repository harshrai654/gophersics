package ch3

import (
	"bytes"
	"fmt"
	"strings"
)

func CommaRecursive(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return CommaRecursive(s[:n-3]) + "," + s[n-3:]
}

func CommaBytesFloat(s string) string {
	if s == "" {
		return s
	}

	var op, num, dec string
	var splits []string

	if s[0] == '+' || s[0] == '-' {
		op = string(s[0])
		splits = strings.Split(s[1:], ".")
	} else {
		splits = strings.Split(s[0:], ".")
	}

	num = splits[0]
	if len(splits) > 1 {
		dec = splits[1]
		var buff bytes.Buffer
		for i := len(dec) - 1; i > -1; i-- {
			fmt.Fprintf(&buff, "%c", dec[i])
		}

		dec = buff.String()
	}

	formattedNum := commaBytes(num)
	formattedDec := commaBytes(dec)

	result := op + formattedNum
	if len(formattedDec) > 0 {
		var buff bytes.Buffer
		for i := len(formattedDec) - 1; i > -1; i-- {
			fmt.Fprintf(&buff, "%c", formattedDec[i])
		}
		result += "." + buff.String()
	}

	return result
}

func commaBytes(s string) string {
	if s == "" {
		return s
	}
	var buf bytes.Buffer
	n := len(s)
	var groupSize int8 = 3

	for i := n - 1; i > 0; i-- {
		buf.WriteString(s[i : i+1])
		groupSize--

		if groupSize == 0 {
			buf.WriteRune(',')
			groupSize = 3
		}
	}
	buf.WriteString(s[0:1])
	reverseBuffer(&buf)

	return buf.String()
}

func reverseBuffer(buf *bytes.Buffer) {
	for i := 0; i < buf.Len()/2; i++ {
		buf.Bytes()[i], buf.Bytes()[buf.Len()-i-1] = buf.Bytes()[buf.Len()-i-1], buf.Bytes()[i]
	}
}

func Anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	freq := make(map[rune]int)

	for _, r := range s1 {
		freq[r]++
	}

	uniqueCount := len(freq)

	for _, r := range s2 {
		freq[r]--
		if freq[r] == 0 {
			uniqueCount--
		}
	}

	return uniqueCount == 0
}
