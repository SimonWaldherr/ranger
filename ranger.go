package ranger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}

func fillWith(integer, count int, char string) int {
	str := "00000000000000" + fmt.Sprint(integer)
	x := str[:len(str)-count] + strings.Repeat(char, count)
	i, _ := strconv.ParseInt(x, 10, 64)
	return int(i)
}

// Compile generates a regular expression that matches any number between the two given numbers
func Compile(min, max int) string {
	var pos []string

	min, max = abs(min), abs(max)

	if min == max {
		return fmt.Sprintf("(%d)", min)
	}

	if min > max {
		min, max = max, min
	}

	if max >= 0 {
		pos = splitPatterns(min, max)
	}

	return strings.Join(pos, "|")
}

func splitPatterns(min, max int) []string {
	var patterns []string

	start := min
	for _, rng := range splitRanges(min, max) {
		if subrange := rangePatterns(start, rng); subrange != "" {
			patterns = append(patterns, subrange)
		}
		start = rng + 1
	}
	return patterns
}

func splitRanges(min, max int) []int {
	stops := []int{max}

	nines := 1
	stop := fillWith(min, nines, "9")
	for min <= stop && stop < max {
		stops = append(stops, stop)
		nines++
		stop = fillWith(min, nines, "9")
	}

	zeros := 1
	stop = fillWith(max+1, zeros, "0") - 1
	for min < stop && stop <= max {
		stops = append(stops, stop)
		zeros++
		stop = fillWith(max+1, zeros, "0") - 1
	}

	sort.Ints(stops)
	return stops
}

func rangePatterns(startint, stopint int) string {
	var pattern string
	var count int
	var startdigit, stopdigit byte

	if startint > stopint {
		startint, stopint = stopint, startint
	}

	start, stop := fmt.Sprint(startint), fmt.Sprint(stopint)

	if startint+1 == stopint && start[0] != stop[0] {
		return ""
	}

	for k := range fmt.Sprint(start) {
		startdigit, stopdigit = start[k]|'0', stop[k]|'0'

		if startdigit == stopdigit || stopdigit == '0' {
			pattern += string(startdigit)
		} else if startdigit != '0' || stopdigit != '9' {
			pattern += fmt.Sprintf("[%c-%c]", startdigit, stopdigit)
		} else {
			count++
		}
	}

	if count == 1 {
		pattern += "\\d"
	} else if count > 1 {
		pattern += fmt.Sprintf("\\d{%d}", count)
	}

	return pattern
}
