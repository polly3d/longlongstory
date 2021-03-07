package main

import (
	"reflect"
	"testing"
)

func Test_shellsort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{[]int{6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6}},
		{"test", args{[]int{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellsort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellsort() = %v, want %v", got, tt.want)
			}
		})
	}
}
