package utils

import (
	"fmt"
	"strings"
)

func ArrayIntegerToString[T int | int64](a []T, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func ArrayStringToString(a []string, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
