package main

import "strings"

func simplifyPath(path string) string {
	parts := strings.FieldsFunc(path, func(r rune) bool {
		return r == '/'
	})

	var stack []string

	for _, part := range parts {
		switch part {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		default:
			stack = append(stack, part)
		}
	}

	res := "/" + strings.Join(stack, "/")

	return res
}
