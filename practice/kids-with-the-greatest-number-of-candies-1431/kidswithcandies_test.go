package main

// The `main` function in a Go program is the entry point for the executable. It is where the
// program starts executing when it is run. In this specific code snippet, the `main` function
// is defining a test function `Test_kidsWithCandies` that tests the `kidsWithCandies` function
// with some sample input and expected output. The `main` function is also defining an `equals`
// function to compare two slices of boolean values for equality.

import "testing"

func equals(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{
			[]int{2, 3, 5, 1, 3},
			3,
			[]bool{true, true, true, false, true},
		},
	}

	for _, test := range tests {
		result := kidsWithCandies(test.candies, test.extraCandies)
		if !equals(result, test.want) {
			t.Errorf("kidsWithCandies(%v, %d) = %v; want %v",
				test.candies, test.extraCandies, result, test.want)
		}
	}
}
