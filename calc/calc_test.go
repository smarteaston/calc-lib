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
