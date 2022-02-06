package assignment

import (
	"math"
	"sort"
	"strings"
)

// AddUint32 adds to untyped int and returns sum of them and overflow
// state of sum
func AddUint32(x, y uint32) (uint32, bool) {

	// sum of them, overflow state
	return x + y, math.MaxUint32-x < y
}

// CeilNumber is rounding numbers to the upper 0.25 bound
func CeilNumber(f float64) float64 {

	var more float64 = 1

	// if remainder is 0, no need to round
	if math.Remainder(f, 0.25) == float64(0) {
		more = 0
	}

	return (math.Floor((f)/0.25) + more) * 0.25
}

// AlphabetSoup is sorting letters in alphabetical order
func AlphabetSoup(s string) string {

	// create array of int version of letters (ascii)
	as_int := []int{}
	for _, v := range s {
		as_int = append(as_int, int(v))
	}

	// sort this numbers
	sort.Ints(as_int)

	// join string back
	result := ""
	for _, v := range as_int {
		result = result + string(rune(v))

	}
	return result
}

// StringMask is masking letters of string after n`th index
func StringMask(s string, n uint) string {

	// special cases -> 0 lenght | 0 number | greater and equal number
	if len(s) == 0 || n == 0 || uint(len(s)) <= n {
		num_of_star := 1

		if len(s) != 0 {
			num_of_star = len(s)
		}

		return strings.Repeat("*", num_of_star)
	}

	// normal case
	return s[:n] + strings.Repeat("*", len(s)-int(n))
}

// WordSplit is splitting strings into possible words
// if contains in a special list
func WordSplit(arr [2]string) string {
	is_possible := false
	word1 := ""
	word2 := ""

Loop:
	for i := 0; i < len(arr[0]); i++ {
		word1 = arr[0][0:i]
		word2 = arr[0][i:]

		// if both words contains in the lists, break the loop
		if strings.Contains(arr[1], word1) && strings.Contains(arr[1], word2) {
			is_possible = true
			break Loop
		}
	}

	// if both words are found
	if is_possible {
		return word1 + "," + word2
	}

	// not found case
	return "not possible"
}

// VariadicSet returns unique elements of list
func VariadicSet(i ...interface{}) []interface{} {

	// map of how many keys there are
	m := make(map[interface{}]int)
	result := []interface{}{}

	for _, v := range i {

		// first time to access value (zero value of int is 0)
		if m[v] == 0 {
			result = append(result, v)
		}

		// increase the value
		m[v] = m[v] + 1
	}

	// return list of unique elements
	return result
}
