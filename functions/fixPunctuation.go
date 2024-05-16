package reload

import (
	"strings"
	"unicode"
)

func FixPunctuationSpacing(input string) string {
	var result strings.Builder

	// Track if the previous character was punctuation
	prevPunctuation := false
	prevSpace := false
	for _, char := range input {

		if unicode.IsSpace(char) { //space
			prevPunctuation = false
			prevSpace = true
			result.WriteRune(char)

		} else if isPunctuation(char) { //punctuation
			if prevPunctuation {
				result.WriteRune(char)
			} else if prevSpace {
				resultString := result.String()
				if len(resultString) > 0 {
					trimmedString := resultString[:len(resultString)-1]
					result.Reset()
					result.WriteString(trimmedString)
				}
				result.WriteRune(char)
			} else {
				result.WriteRune(char)
			}
			prevPunctuation = true
			prevSpace = false

		} else { //char
			if prevPunctuation {
				result.WriteRune(' ')
				result.WriteRune(char)
			} else {
				result.WriteRune(char)
			}
			prevPunctuation = false
			prevSpace = false

		}

	}

	return result.String()
}
