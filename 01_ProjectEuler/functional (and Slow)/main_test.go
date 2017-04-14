package main

import "testing"

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

type testpair struct {
	number     int
	multipleOf int
	expected   bool
}

var tests = []testpair{
	{0, 3, false},
	{1, 3, false},
	{2, 3, false},
	{3, 3, true},
	{4, 3, false},
	{6, 3, true},
}

func TestMultipleOf(t *testing.T) {
	for _, pair := range tests {
		v := isMultipleOf(pair.number, pair.multipleOf)
		if v != pair.expected {
			t.Error(
				"For", pair.number,
				"With", pair.multipleOf,
				"expected", pair.expected,
				"got", v,
			)
		}
	}
}

func TestItemInSlice(t *testing.T) {
	testSlice := []int{1, 4, 5, 12, 20, 30}

	v := isInSliceInt(4, testSlice)
	if !v {
		t.Error("For 4, Expected True, Got:", v)
	}

	v = isInSliceInt(30, testSlice)
	if !v {
		t.Error("For 4, Expected True, Got:", v)
	}

	v = isInSliceInt(3, testSlice)
	if v {
		t.Error("For 4, Expected True, Got:", v)
	}
}

func TestMultiplesOfNumberBelow(t *testing.T) {
	var actual []int
	expected := []int{3, 5, 6, 9}
	actual = multiplesOfNumberBelow(10, 3, 5)

	for i, v := range expected {
		if v != actual[i] {
			t.Error("Expected", expected[i], "Got:", v)
		}
	}
}

func TestAddNumbers(t *testing.T) {
	testSlice := []int{3, 5, 6, 9}
	expected := 23

	v := addNumbers(testSlice...)
	if v != expected {
		t.Error("Expected 36, Got:", v)
	}
}

func TestProgram(t *testing.T) {
	expected := 23
	v := addNumbers(multiplesOfNumberBelow(10, 3, 5)...)

	if v != expected {
		t.Error("Expected 23, Got:", v)
	}
}

func BenchmarkProgram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addNumbers(multiplesOfNumberBelow(i, 3, 5)...)
	}
}
