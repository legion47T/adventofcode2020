package main

import (
	"testing"
)

func Test_findSeat(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want seat
	}{
		{
			"BFFFBBFRRR is row 70, column 7, seat ID 567",
			args{"BFFFBBFRRR"},
			seat{
				code:   codeT{"BFFFBBFRRR", "BFFFBBF", "RRR"},
				row:    70,
				column: 7,
				id:     567,
			},
		},
		{
			"FFFBBBFRRR is row 14, column 7, seat ID 119",
			args{"FFFBBBFRRR"},
			seat{
				code:   codeT{"FFFBBBFRRR", "FFFBBBF", "RRR"},
				row:    14,
				column: 7,
				id:     119,
			},
		},
		{
			"BBFFBBFRLL is row 102, column 4, seat ID 820",
			args{"BBFFBBFRLL"},
			seat{
				code:   codeT{"BBFFBBFRLL", "BBFFBBF", "RLL"},
				row:    102,
				column: 4,
				id:     820,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSeat(tt.args.code); got != tt.want {
				t.Errorf("findSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRow(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"FBFBBFF is 44",
			args{"FBFBBFF"},
			44,
		},
		{
			"BFFFBBF is 70",
			args{"BFFFBBF"},
			70,
		},
		{
			"FFFBBBF is 14",
			args{"FFFBBBF"},
			14,
		},
		{
			"BBFFBBF is 102",
			args{"BBFFBBF"},
			102,
		},
		{
			"RLR is 5",
			args{"RLR"},
			5,
		},
		{
			"RRR is 7",
			args{"RRR"},
			7,
		},
		{
			"RLL is 4",
			args{"RLL"},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRowOrColumn(tt.args.code); got != tt.want {
				t.Errorf("calculateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want codeT
	}{
		{
			"BFFFBBFRRR is split to BFFFBBF and RRR",
			args{"BFFFBBFRRR"},
			codeT{
				full:   "BFFFBBFRRR",
				row:    "BFFFBBF",
				column: "RRR",
			},
		},
		{
			"FFFBBBFRRR is split to FFFBBBF and RRR",
			args{"FFFBBBFRRR"},
			codeT{
				full:   "FFFBBBFRRR",
				row:    "FFFBBBF",
				column: "RRR",
			},
		},
		{
			"BBFFBBFRLL is split to BBFFBBF and RLL",
			args{"BBFFBBFRLL"},
			codeT{
				full:   "BBFFBBFRLL",
				row:    "BBFFBBF",
				column: "RLL",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitCode(tt.args.code); got != tt.want {
				t.Errorf("calculateID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateID(t *testing.T) {
	type args struct {
		row    int
		column int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"row 70, column 7 is ID 567",
			args{70, 7},
			567,
		},
		{
			"row 14, column 7 is ID 119",
			args{14, 7},
			119,
		},
		{
			"row 102, column 4 is ID 820",
			args{102, 4},
			820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateID(tt.args.row, tt.args.column); got != tt.want {
				t.Errorf("calculateID() = %v, want %v", got, tt.want)
			}
		})
	}
}
