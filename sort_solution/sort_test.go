package sort_solution

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: []int{8, 9, 7, 4, 2, 1, 3, 5, 6, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shellSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{arr: []int{8, 9, 7, 4, 2, 1, 3, 5, 6, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shellSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shellSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{8, 9, 7, 4, 2, 1, 3, 5, 6, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_quickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{8, 9, 7, 4, 2, 1, 3, 5, 6, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("quickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{nums: []int{8, 9, 7, 4, 2, 1, 3, 5, 6, 10}},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
