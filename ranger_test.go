package ranger

import (
	"fmt"
	"regexp"
	"testing"

	"simonwaldherr.de/go/golibs/as"
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

func BenchmarkNumberRegEx(b *testing.B) {
	re := Compile(89, 1001)
	re = "^("+re+")$"
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		matched, err := regexp.MatchString(re, "404")
		if !matched || err != nil {
			b.Log("Error in Benchmark")
		}
		
		matched, err = regexp.MatchString(re, "2000")
		if matched || err != nil {
			b.Log("Error in Benchmark")
		}
	}
}

func BenchmarkFulltextRegEx(b *testing.B) {
	re := Compile(89, 1001)
	re = " ("+re+") "
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		matched, err := regexp.MatchString(re, "lorem ipsum 404 dolor sit")
		if !matched || err != nil {
			b.Log("Error in Benchmark")
		}

		matched, err = regexp.MatchString(re, "lorem ipsum 2000 dolor sit")
		if matched || err != nil {
			b.Log("Error in Benchmark")
		}
	}
}

func BenchmarkNumberParse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		i1 := as.Int("404")
		i2 := as.Int("2000")
		
		if i1 < 89 || i1 > 1001 {
			b.Log("Error in Benchmark")
		}

		if !(i2 < 89 || i2 > 1001) {
			b.Log("Error in Benchmark")
		}
	}
}

func BenchmarkFulltextParse(b *testing.B) {
	re := regexp.MustCompile("[0-9]+")
	b.ResetTimer()
	
	for n := 0; n < b.N; n++ {
		i1 := as.Int(re.FindString("lorem ipsum 404 dolor sit"))
		i2 := as.Int(re.FindString("lorem ipsum 2000 dolor sit"))

		if i1 < 89 || i1 > 1001 {
			b.Log("Error in Benchmark")
		}

		if !(i2 < 89 || i2 > 1001) {
			b.Log("Error in Benchmark")
		}
	}
}
