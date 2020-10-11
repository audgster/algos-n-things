package algos

import "testing"

type permutationTestCase struct {
	a        string
	b        string
	expected bool
	testcase string
}

func TestIsPermutation(t *testing.T) {
	testcases := []permutationTestCase{
		{
			a:        "abc",
			b:        "abc",
			expected: true,
			testcase: "Both strings are identical",
		},
		{
			a:        "abc",
			b:        "bca",
			expected: true,
			testcase: "Out of order permutation",
		},
		{
			a:        "abc",
			b:        "bcd",
			expected: false,
			testcase: "Same length string, different charset",
		},
		{
			a:        "abc",
			b:        "aabbcc",
			expected: false,
			testcase: "strings are not the same length",
		},
	}

	for _, test := range testcases {
		actual := IsPermutation(test.a, test.b)

		if actual != test.expected {
			t.Logf("For testcase %s with strings %s and %s: expected %t, got %t", test.testcase, test.a, test.b, test.expected, actual)
			t.Fail()
		}
	}
}

type replaceSpacesTestCase struct {
	input      string
	truelength int
	expected   string
	testcase   string
}

var replaceSpacesTestcases []replaceSpacesTestCase = []replaceSpacesTestCase{
	{
		input:      "My String  ",
		truelength: 9,
		expected:   "My%20String",
		testcase:   "Single space with trailing spaces",
	},
	{
		input:      "My String meep    ",
		truelength: 14,
		expected:   "My%20String%20meep",
		testcase:   "Multiple spaces with trailing spaces",
	},
	{
		input:      "nospaces",
		truelength: 8,
		expected:   "nospaces",
		testcase:   "No spaces",
	},
	{
		input:      "         ",
		truelength: 3,
		expected:   "%20%20%20",
		testcase:   "All spaces",
	},
	{
		input:      "",
		truelength: 0,
		expected:   "",
		testcase:   "Empty String",
	},
}

func TestReplaceSpacesBuffer(t *testing.T) {
	for _, test := range replaceSpacesTestcases {
		actual := ReplaceSpacesBuffer(test.input, test.truelength)

		if actual != test.expected {
			t.Logf("For testcase %s with input %q with true length %d: expected %q, got %q", test.testcase, test.input, test.truelength, test.expected, actual)
			t.Fail()
		}
	}
}

func TestReplaceSpacesInPlace(t *testing.T) {
	for _, test := range replaceSpacesTestcases {
		actual := ReplaceSpacesInPlace([]rune(test.input), test.truelength)

		if actual != test.expected {
			t.Logf("For testcase %s with input %q with true length %d: expected %q, got %q", test.testcase, test.input, test.truelength, test.expected, actual)
			t.Fail()
		}
	}
}
