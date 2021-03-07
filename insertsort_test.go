package main

import (
	"reflect"
	"testing"
)

func Test_insertsort1(t *testing.T) {
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
			if got := insertsort1(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertsort1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertsort2(t *testing.T) {
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
			if got := insertsort2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertsort2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_insertsort1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertsort1([]int{6, 5, 4, 3, 2, 1})
	}
}

func Benchmark_insertsort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertsort2([]int{6, 5, 4, 3, 2, 1})
	}
}
