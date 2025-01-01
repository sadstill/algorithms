package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 8, 4, 2}, expected: []int{2, 3, 4, 5, 8}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: []int{42}, expected: []int{42}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("TestCase #%d", i+1), func(t *testing.T) {
			result := make([]int, len(test.input))
			copy(result, test.input)
			quickSort(result)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Тест-кейс #%d: Для входных данных %v ожидалось %v, но получено %v", i+1, test.input, test.expected, result)
			}
		})
	}
}
