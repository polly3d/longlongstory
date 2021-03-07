package main

func selectedsort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[i] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func selectedSortMax(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		max := i
		for j := i; j >= 0; j-- {
			if arr[j] > arr[i] {
				max = j
			}
		}
		arr[max], arr[i] = arr[i], arr[max]
	}
	return arr
}
