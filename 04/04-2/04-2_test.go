package main

import (
	"testing"
)

func Test_isPidValid(t *testing.T) {
	type args struct {
		pid string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"000000001 is valid",
			args{"000000001"},
			true,
		},
		{
			"087499704 is valid",
			args{"087499704"},
			true,
		},
		{
			"0123456789 is invalid",
			args{"0123456789"},
			false,
		},
		{
			"186cm is invalid",
			args{"186cm"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPidValid(tt.args.pid); got != tt.want {
				t.Errorf("isPidValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEyeClrValid(t *testing.T) {
	type args struct {
		inClr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"brn is valid",
			args{"brn"},
			true,
		},
		{
			"wat is invalid",
			args{"wat"},
			false,
		},
		{
			"zzz is invalid",
			args{"zzz"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEyeClrValid(tt.args.inClr); got != tt.want {
				t.Errorf("isEyeClrValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHairClrValid(t *testing.T) {
	type args struct {
		clr string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"#123abc is valid",
			args{"#123abc"},
			true,
		},
		{
			"#623a2f is valid",
			args{"#623a2f"},
			true,
		},
		{
			"#123abz is invalid",
			args{"#123abz"},
			false,
		},
		{
			"123abz is invalid",
			args{"123abz"},
			false,
		},
		{
			"dab227 is invalid",
			args{"dab227"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHairClrValid(tt.args.clr); got != tt.want {
				t.Errorf("isHairClrValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHeightValid(t *testing.T) {
	type args struct {
		heigth string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"60in is valid",
			args{"60in"},
			true,
		},
		{
			"190cm is valid",
			args{"190cm"},
			true,
		},
		{
			"165cm is valid",
			args{"165cm"},
			true,
		},
		{
			"190in is invalid",
			args{"190in"},
			false,
		},
		{
			"190 is invalid",
			args{"190"},
			false,
		},
		{
			"59cm is invalid",
			args{"59cm"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHeightValid(tt.args.heigth); got != tt.want {
				t.Errorf("isHeightValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
