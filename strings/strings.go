package strings

import "strings"

func IsBlank(input string) bool {
	input = ReplaceNewLineAndTabToSpace(input)
	return IsEmpty(input)
}

func IsEmpty(input string) bool {
	return len(input) == 0
}

func ReplaceNewLineAndTabToSpace(input string) string {
	result := strings.ReplaceAll(input, "\n", " ")
	result = strings.ReplaceAll(result, "\r", " ")
	result = strings.ReplaceAll(result, "\t", " ")
	result = strings.TrimSpace(result)
	return result
}
