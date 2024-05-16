package solution

func Solution(A int, B int) int {
	countPieces := func(length int, stick1 int, stick2 int) int {
		return (stick1 / length) + (stick2 / length)
	}

	canFormSquare := func(length int, stick1 int, stick2 int) bool {
		return countPieces(length, stick1, stick2) >= 4
	}

	var start int
	if A < B {
		start = A
	} else {
		start = B
	}

	for length := start; length > 0; length-- {
		if canFormSquare(length, A, B) {
			return length
		}
	}
	return 0
}
