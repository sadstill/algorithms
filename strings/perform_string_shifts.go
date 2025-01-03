package main

func stringShift(s string, shift [][]int) string {
	netShift := 0

	for _, sh := range shift {
		direction, amount := sh[0], sh[1]
		if direction == 0 {
			netShift -= amount
		} else {
			netShift += amount
		}
	}

	netShift = ((netShift % len(s)) + len(s)) % len(s)

	if netShift == 0 {
		return s
	}

	return s[len(s)-netShift:] + s[:len(s)-netShift]
}
