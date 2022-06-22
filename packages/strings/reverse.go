package strings

func Reverse(s string) string {
	runes := []rune(s)
	reverseRunes := reverseRunes(runes)
	return string(reverseRunes)
}
