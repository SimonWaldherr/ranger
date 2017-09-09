package ranger_test

import (
	"fmt"
	"regexp"
	"simonwaldherr.de/go/ranger"
)

func ExampleCompile() {
	re := ranger.Compile(23, 37)
	fmt.Println("^(" + re + ")$")

	for v := range [200]int{} {
		matched, err := regexp.MatchString("^("+re+")$", fmt.Sprint(v))
		if err != nil {
			fmt.Println(matched, err)
		}
		if matched {
			fmt.Print(v, " ")
		}
	}
	// Output: ^(2[3-9]|3[0-7])$
	// 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37
}
