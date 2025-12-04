package numbers

func Length(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	length := 0
	for n > 0 {
		n /= 10
		length++
	}

	return length
}
