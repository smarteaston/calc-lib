package calc

import (
	"fmt"
	"math"
	"testing"
)

func TestAddition_Calculate(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{a: 0, b: 0, want: 0},
		{a: 0, b: 1, want: 1},
		{a: 1, b: 0, want: 1},
		{a: 1, b: 1, want: 2},
		{a: 2, b: 3, want: 5},
		{a: -1, b: 1, want: 0},
		{a: math.MaxInt, b: 1, want: math.MinInt},
		{a: 3, b: 1, want: 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d + %d = %d", tt.a, tt.b, tt.want), func(t *testing.T) {
			this := Addition{}
			if got := this.Calculate(tt.a, tt.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{}, 0},
		{"positive", args{5, 3}, 2},
		{"negative", args{3, 7}, -4},
		{"sub zero", args{3, 0}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Subtract{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty", args{}, 0},
		{"positive", args{1, 3}, 3},
		{"negative", args{-1, 3}, -3},
		{"zero", args{5, 0}, 0},
		{"normal", args{5, 4}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Multiply{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"empty almost", args{0, 1}, 0},
		{"no remainder", args{25, 5}, 5},
		{"integer division", args{25, 4}, 6},
		{"negative", args{4, -3}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Division{}
			if got := this.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivision_ByZeroPanics(t *testing.T) {
	division := &Division{}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("Division.ByZero did not panic")
		}
	}()
	division.Calculate(1, 0)
}
