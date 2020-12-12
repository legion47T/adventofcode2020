package main

import (
	"testing"
)

func Test_willBeOccupied(t *testing.T) {
	emptyRoom := [][]seat{
		{{0, 0, true, false}, {1, 0, true, false}, {2, 0, true, false}},
		{{0, 1, true, false}, {1, 1, true, false}, {2, 1, true, false}},
		{{0, 2, true, false}, {1, 2, true, false}, {2, 2, true, false}}}
	fourTakenSeatsRoom := [][]seat{
		{{0, 0, true, false}, {1, 0, true, true}, {2, 0, true, false}},
		{{0, 1, true, true}, {1, 1, true, false}, {2, 1, true, true}},
		{{0, 2, true, false}, {1, 2, true, true}, {2, 2, true, false}}}

	type args struct {
		room *[][]seat
		s    seat
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Seat 0,0 will be occupied if no other seat is taken",
			args{
				&emptyRoom,
				seat{0, 0, true, false},
			},
			true,
		},
		{
			"Seat 1,1 will be occupied if no other seat is taken",
			args{
				&emptyRoom,
				seat{1, 1, true, false},
			},
			true,
		},
		{
			"Seat 2,2 will be occupied if no other seat is taken",
			args{
				&emptyRoom,
				seat{2, 2, true, false},
			},
			true,
		},
		{
			"Seat 0,0 will be occupied if two surrounding seats are taken",
			args{
				&fourTakenSeatsRoom,
				seat{0, 0, true, false},
			},
			true,
		},
		{
			"Seat 1,1 will not be occupied if four surrounding seats are taken",
			args{
				&fourTakenSeatsRoom,
				seat{1, 1, true, false},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := willBeOccupied(tt.args.room, tt.args.s); got != tt.want {
				t.Errorf("willBeOccupied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stateHasChanged(t *testing.T) {
	emptyRoom := [][]seat{
		{{0, 0, true, false}, {1, 0, true, false}, {2, 0, true, false}},
		{{0, 1, true, false}, {1, 1, true, false}, {2, 1, true, false}},
		{{0, 2, true, false}, {1, 2, true, false}, {2, 2, true, false}}}
	fourTakenSeatsRoom := [][]seat{
		{{0, 0, true, false}, {1, 0, true, true}, {2, 0, true, false}},
		{{0, 1, true, true}, {1, 1, true, false}, {2, 1, true, true}},
		{{0, 2, true, false}, {1, 2, true, true}, {2, 2, true, false}}}
	type args struct {
		currentRoom *[][]seat
		futureRoom  *[][]seat
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"if two rooms are equal state has not changed",
			args{
				&emptyRoom,
				&emptyRoom,
			},
			false,
		},
		{
			"if two rooms are unequal state has changed",
			args{
				&emptyRoom,
				&fourTakenSeatsRoom,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stateHasChanged(tt.args.currentRoom, tt.args.futureRoom); got != tt.want {
				t.Errorf("stateHasChanged() = %v, want %v", got, tt.want)
			}
		})
	}
}
