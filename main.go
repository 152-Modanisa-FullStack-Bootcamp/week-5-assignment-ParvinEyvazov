package main

import (
	"bootcamp/assignment"
	"fmt"
)

func main() {
	words := "apple,bat,cat,goodbye,hello,yellow,why"

	cases := []struct {
		given    string
		expected string
	}{
		{"hellocat", "hello,cat"},
		{"catbat", "cat,bat"},
		{"yellowapple", "yellow,apple"},
		{"", "not possible"},
		{"notcat", "not possible"},
		{"bootcamprocks!", "not possible"},
	}

	for _, v := range cases {
		fmt.Println(assignment.WordSplit([2]string{v.given, words}))
	}
}
