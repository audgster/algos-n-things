package algos

import "testing"

type permutationTestCase struct {
	a string
	b string
	expected bool
}

func TestIsPermutation(t *testing.T) {
	testcases := []permutationTestCase{
		{
			a: "abc",
			b: "abc",
			expected
		}, // same string are permutations
		"abc": "bca", // actual permutation
	}
}
