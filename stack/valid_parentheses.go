package main

func isValid(s string) bool {
	pairs := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	var stack []string

	for _, c := range s {
		if _, ok := pairs[string(c)]; ok {
			stack = append(stack, string(c))
		} else {
			if len(stack) == 0 {
				return false
			}

			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pairs[prev] != string(c) {
				return false
			}
		}
	}

	return len(stack) == 0
}
