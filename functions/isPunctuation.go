package reload

func isPunctuation(char rune) bool {
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}
	for _, p := range punctuations {
		if char == p {
			return true
		}
	}
	return false
}
