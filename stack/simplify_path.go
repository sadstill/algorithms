package main

import "strings"

func simplifyPath(path string) string {
	parts := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	stack := make([]string, 0, len(parts))

	for _, part := range parts {
		if part == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				continue
			}
		} else if part == "." {
			continue
		} else {
			stack = append(stack, part)
		}
	}

	res := "/" + strings.Join(stack, "/")

	return res
}
