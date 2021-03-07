package main

import "testing"

func TestSortCompare(t *testing.T) {
	tests := []struct {
		name     string
		sortType string
		nums     int
	}{
		{"test", "insertsort", 1000},
		{"test", "selectedsort", 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortCompare(tt.sortType, tt.nums)
		})
	}
}
