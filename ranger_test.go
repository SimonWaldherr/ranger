package ranger

import (
	"fmt"
	"regexp"
	"testing"
)

type test struct {
	min int
	max int
}

var tests = []test{{
	min: -4,
	max: 14,
}, {
	min: 0,
	max: 24,
}, {
	min: 10,
	max: 10,
}, {
	min: 18,
	max: 10,
}, {
	min: 23,
	max: 37,
}, {
	min: 11,
	max: 42,
}, {
	min: 2,
	max: 200,
}, {
	min: 42,
	max: 1337,
}}

func Test_Compile(t *testing.T) {
	for _, tt := range tests {
		min, max := tt.min, tt.max
		
		re := Compile(min, max)
		
		min, max = abs(min), abs(max)
		
		if min > max {
			min, max = max, min
		}

		for v := range [2000]int{} {
			matched, err := regexp.MatchString("^("+re+")$", fmt.Sprint(v))
			if err != nil {
				t.Fatalf("Compile Test failed: %v\n", err)
			}
			if matched && (v < min || v > max) {
				t.Fatalf("Compile Test failed: %v | re: %v | i: %v\n", "false positive", re, v)
			} else if !matched && v >= min && v <= max {
				t.Fatalf("Compile Test failed: %v | re: %v | i: %v\n", "false negative", re, v)
			}
		}
	}
}
