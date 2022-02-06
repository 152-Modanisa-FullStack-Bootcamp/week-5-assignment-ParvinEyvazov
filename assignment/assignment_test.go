package assignment

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUint32(t *testing.T) {

	cases := []struct {
		given_numbers     [2]uint32
		expected_result   uint32
		expected_overflow bool
	}{
		{[2]uint32{math.MaxUint32, 1}, 0, true},
		{[2]uint32{1, 1}, 2, false},
		{[2]uint32{42, 2701}, 2743, false},
		{[2]uint32{42, math.MaxUint32}, 41, true},
		{[2]uint32{4294967290, 5}, 4294967295, false},
		{[2]uint32{4294967290, 6}, 0, true},
		{[2]uint32{4294967290, 10}, 4, true},
	}

	for _, value := range cases {
		t.Run(
			fmt.Sprintf("%v+%v", value.given_numbers[0], value.given_numbers[1]),
			func(t *testing.T) {
				result, overflow := AddUint32(value.given_numbers[0], value.given_numbers[1])
				assert.Equal(t, value.expected_result, result)
				assert.Equal(t, value.expected_overflow, overflow)
			},
		)
	}
}

func TestCeilNumber(t *testing.T) {

	cases := []struct {
		given    float64
		expected float64
	}{
		{42.42, 42.50}, {42, 42}, {42.01, 42.25},
		{42.24, 42.25}, {42.25, 42.25}, {42.26, 42.50},
		{42.55, 42.75}, {42.75, 42.75}, {42.76, 43},
		{42.99, 43}, {43.13, 43.25},
	}

	for _, value := range cases {
		t.Run(fmt.Sprintf("%v=>%v", value.given, value.expected),
			func(t *testing.T) {
				result := CeilNumber(value.given)
				assert.Equal(t, value.expected, result)
			},
		)
	}
}

func TestAlphabetSoup(t *testing.T) {

	cases := []struct {
		given    string
		expected string
	}{
		{"hello", "ehllo"}, {"", ""},
		{"h", "h"}, {"ab", "ab"},
		{"ba", "ab"}, {"bac", "abc"},
		{"cba", "abc"},
	}

	fmt.Println(cases)

	for _, v := range cases {
		t.Run(
			v.given,
			func(t *testing.T) {
				result := AlphabetSoup(v.given)
				assert.Equal(t, v.expected, result)
			})
	}
}

func TestStringMask(t *testing.T) {

	cases := []struct {
		word     string
		len      uint
		expected string
	}{
		{"!mysecret*", 2, "!m********"},
		{"", 8, "*"},
		{"a", 1, "*"},
		{"string", 0, "******"},
		{"string", 3, "str***"},
		{"string", 5, "strin*"},
		{"string", 6, "******"},
		{"string", 7, "******"},
		{"s*r*n*", 3, "s*r***"},
	}

	for _, v := range cases {
		t.Run(
			fmt.Sprintf("%v => %v", v.word, v.len),
			func(t *testing.T) {
				result := StringMask(v.word, v.len)
				assert.Equal(t, v.expected, result)
			})
	}
}

func TestWordSplit(t *testing.T) {
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
		t.Run(
			fmt.Sprintf("%v", v),
			func(t *testing.T) {
				result := WordSplit([2]string{v.given, words})
				assert.Equal(t, v.expected, result)
			})
	}
}

func TestVariadicSet(t *testing.T) {

	cases := []interface{}{
		[]interface{}{4, 2, 5, 4, 2, 4},
		[]interface{}{"bootcamp", "rocks!", "really", "rocks! "},
		[]interface{}{1, uint32(1), "first", 2, uint32(2), "second", 1, uint32(2), "first"},
	}

	fmt.Println(cases)

	/*
		FINAL BOSS ALERT :)
		Tip: Learn and apply golang variadic functions(search engine -> "golang variadic function" -> WOW You can really dance! )

		Convert inputs to set(no duplicate element)
		cases need to pass:
			4,2,5,4,2,4 => []interface{4,2,5}
			"bootcamp","rocks!","really","rocks! => []interface{"bootcamp","rocks!","really"}
			1,uint32(1),"first",2,uint32(2),"second",1,uint32(2),"first" => []interface{1,uint32(1),"first",2,uint32(2),"second"}
	*/
	set := VariadicSet(4, 2, 5, 4, 2, 4)

	assert.Equal(t, []interface{}{4, 2, 5}, set)
}
