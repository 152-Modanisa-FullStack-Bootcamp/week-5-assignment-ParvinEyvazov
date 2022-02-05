package assignment

import (
	"math"
)

func AddUint32(x, y uint32) (uint32, bool) {
	return x + y, math.MaxUint32-x < y
}

func CeilNumber(f float64) float64 {

	if math.Remainder(f, 0.25) == float64(0) {
		return math.Floor(f/0.25) * 0.25
	}

	return (math.Floor((f)/0.25) + 1) * 0.25
}

func AlphabetSoup(s string) string {
	return ""
}

func StringMask(s string, n uint) string {
	return ""
}

func WordSplit(arr [2]string) string {
	return ""
}

func VariadicSet(i ...interface{}) []interface{} {
	return nil
}
