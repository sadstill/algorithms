package main

func main() {
	s := "abc"
	t := "ahbgdc"
	isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	l := 0
	r := 0

	for r < len(t) {
		if s[l] == t[r] {
			l++
		}
		r++
		if l == len(s) {
			return true
		}
	}

	return false
}
