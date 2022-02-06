package assignment

import (
	"math"
	"sort"
	"strings"
)

func AddUint32(x, y uint32) (uint32, bool) {
	return x + y, math.MaxUint32-x < y
}

func CeilNumber(f float64) float64 {

	var more float64 = 1

	if math.Remainder(f, 0.25) == float64(0) {
		more = 0
	}

	return (math.Floor((f)/0.25) + more) * 0.25
}

func AlphabetSoup(s string) string {
	as_int := []int{}
	for _, v := range s {
		as_int = append(as_int, int(v))
	}
	sort.Ints(as_int)

	result := ""
	for _, v := range as_int {
		result = result + string(rune(v))

	}
	return result
}

func StringMask(s string, n uint) string {
	if len(s) == 0 || n == 0 || uint(len(s)) <= n {
		num_of_star := 1

		if len(s) != 0 {
			num_of_star = len(s)
		}

		return strings.Repeat("*", num_of_star)
	}

	return s[:n] + strings.Repeat("*", len(s)-int(n))
}

func WordSplit(arr [2]string) string {
	is_possible := false
	word1 := ""
	word2 := ""

Loop:
	for i := 0; i < len(arr[0]); i++ {
		word1 = arr[0][0:i]
		word2 = arr[0][i:]

		if strings.Contains(arr[1], word1) && strings.Contains(arr[1], word2) {
			is_possible = true
			break Loop
		}
	}

	if is_possible {
		return word1 + "," + word2
	}

	return "not possible"
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
