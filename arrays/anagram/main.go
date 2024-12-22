package main

import "fmt"

func main() {
	//s := "пello"
	//fmt.Println(string([]rune(s)[0]))

	fmt.Println(byte('a'), byte('b'))

}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap, tMap := make(map[rune]int), make(map[rune]int)
	sRuneArr, tRuneArr := []rune(s), []rune(t)

	for i := range len([]rune(s)) {
		sMap[sRuneArr[i]]++
		tMap[tRuneArr[i]]++
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}
