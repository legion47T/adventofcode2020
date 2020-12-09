package main

import "testing"

func Test_sumUp(t *testing.T) {
	testNums := []int{1, 3, 9, 12}
	type args struct {
		nums *[]int
		lo   int
		hi   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sum 4 when lo 0, hi 1",
			args{
				nums: &testNums,
				lo:   0,
				hi:   1,
			},
			4,
		},
		{
			"sum 13 when lo 0, hi 2",
			args{
				nums: &testNums,
				lo:   0,
				hi:   2,
			},
			13,
		},
		{
			"sum 21 when lo 2, hi 3",
			args{
				nums: &testNums,
				lo:   2,
				hi:   3,
			},
			21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumUp(tt.args.nums, tt.args.lo, tt.args.hi); got != tt.want {
				t.Errorf("sumUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
