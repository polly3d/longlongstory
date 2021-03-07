package main

import (
	"reflect"
	"testing"
)

func Test_selectedsort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{[]int{6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectedsort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectedsort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_selectedSortMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test", args{[]int{6, 5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := selectedSortMax(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selectedSortMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
