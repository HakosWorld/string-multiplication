package reload

func IsAlpha(s string) bool {
	for _, c := range s {
		if (c < 'a' || c > 'z') || (c < 'A' || c > 'Z') {
			return false
		}
	}
	return true
}
