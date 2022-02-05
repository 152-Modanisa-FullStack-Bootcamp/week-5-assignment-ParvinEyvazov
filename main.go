package main

import (
	"bootcamp/assignment"
	"fmt"
	"math"
)

func main() {

	result, overflow := assignment.AddUint32(42, math.MaxUint32)

	fmt.Printf(`result %v, overflow %v`, result, overflow)

}
