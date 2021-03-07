package main

func insertsort1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

//此方法要慢于上一种，可能在go上比较特殊
func insertsort2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		t := arr[i]
		j := i
		for ; j > 0 && t < arr[j-1]; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = t
	}
	return arr
}
