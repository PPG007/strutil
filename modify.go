package strutil

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Reverse reverses the string
// Copied from here https://stackoverflow.com/a/20225618/153570
func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

// ReplaceAllToOne replaces every string in the from to the string "to"
func ReplaceAllToOne(str string, from []string, to string) string {
	arr := make([]string, len(from)*2)
	for i, s := range from {
		arr[i*2] = s
		arr[i*2+1] = to
	}
	r := strings.NewReplacer(arr...)

	return r.Replace(str)
}

// MapLines runs function fn on every line of the string.
// It splits the string by new line "\n" and runs the fn for every line and
// returns the new string by combining these lines with "\n"
func MapLines(str string, fn func(string) string) string {
	arr := strings.Split(str, "\n")
	for i := 0; i < len(arr); i++ {
		arr[i] = fn(arr[i])
	}
	return strings.Join(arr, "\n")
}

// Splice insert a new string in place of the string between start and end indexes.Splice
// It is based on runes so start and end indexes are rune based indexes.
// It can be used to remove a part of string by giving newStr as empty string
func Splice(str string, newStr string, start int, end int) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	size := len(runes)
	if start < 0 || start > size-1 {
		panic(fmt.Sprintf("start (%d) is out of range (%d)", start, size))
	}
	if end <= start || end > size {
		panic(fmt.Sprintf("end (%d) is out of range (%d)", end, size))
	}
	return string(runes[:start]) + newStr + string(runes[end:])
}
