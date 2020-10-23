package strutil

import "strings"

// Reverse is a function that reverses a string
func Reverse(s string) string {
	stringList := strings.Split(s, "")
	for i, j := 0, len(stringList)-1; i < j; i, j = i+1, j-1 {
		stringList[i], stringList[j] = stringList[j], stringList[i]
	}
	return strings.Join(stringList, "")
}
