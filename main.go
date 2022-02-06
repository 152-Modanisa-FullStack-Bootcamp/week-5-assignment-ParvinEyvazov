package main

import (
	"bootcamp/assignment"
	"fmt"
)

func main() {
	cases := []struct {
		given    []interface{}
		expected []interface{}
	}{
		{
			[]interface{}{4, 2, 5, 4, 2, 4},
			[]interface{}{4, 2, 5},
		},
		{
			[]interface{}{"bootcamp", "rocks!", "really", "rocks!"},
			[]interface{}{"bootcamp", "rocks!", "really"},
		},
		{
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
			[]interface{}{1, uint32(1), "first", 2, uint32(2), "second"},
		},
	}

	result := assignment.VariadicSet(cases[2].given...)

	for _, v := range result {
		fmt.Printf("%v %T\n", v, v)
	}

	/* 	for _, v := range cases {
		fmt.Println(assignment.VariadicSet(v))
	} */
}
