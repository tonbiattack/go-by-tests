package stringexample

import "unicode/utf8"

func ByteLen(value string) int {
	return len(value)
}

func RuneLen(value string) int {
	return utf8.RuneCountInString(value)
}
