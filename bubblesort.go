package main

func bubblesort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}
