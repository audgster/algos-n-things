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

type stringCompressionTestCase struct {
	input    string
	expected string
	testcase string
}

func TestStringCompression(t *testing.T) {

	testcases := []stringCompressionTestCase{
		{
			input:    "aaabbbcccdeef",
			expected: "a3b3c3d1e2f1",
			testcase: "standard compression",
		},
		{
			input:    "abbccddd",
			expected: "a1b2c2d3",
			testcase: "first caracter is alone",
		},
		{
			input:    "aabbac",
			expected: "a2b2a1c1",
			testcase: "repeated charsets",
		},
		{
			input:    "nodups",
			expected: "nodups",
			testcase: "uncompressable string returns itself",
		},
		{
			input:    "a",
			expected: "a",
			testcase: "edge: single char",
		},
		{
			input:    "",
			expected: "",
			testcase: "empty string",
		},
	}

	for _, test := range testcases {
		actual := StringCompression(test.input)

		if actual != test.expected {
			t.Logf("For testcase %s with input %q expected %q, got %q", test.testcase, test.input, test.expected, actual)
			t.Fail()
		}
	}

}

type zeroMatrixRowColumnTestCase struct {
	input    [][]int
	expected [][]int
	testcase string
}

func TestZeroMatrixRowColumn(t *testing.T) {
	testcases := []zeroMatrixRowColumnTestCase{
		{
			input: [][]int{
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
			},
			expected: [][]int{
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
			},
			testcase: "no zero bits",
		},
		{
			input: [][]int{
				[]int{0, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
				[]int{1, 1, 1, 1},
			},
			expected: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 1, 1, 1},
				[]int{0, 1, 1, 1},
				[]int{0, 1, 1, 1},
			},
			testcase: "one zero",
		},
		{
			input: [][]int{
				[]int{0, 1, 1, 1},
				[]int{1, 0, 1, 1},
				[]int{1, 1, 0, 1},
				[]int{1, 1, 1, 0},
			},
			expected: [][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
				[]int{0, 0, 0, 0},
			},
			testcase: "diagonal zero",
		},
	}

	for _, test := range testcases {
		actual := ZeroMatrixRowColumn(test.input)

		if !areEqual(actual, test.expected) {
			t.Logf("For testcase %s with input %q expected %q, got %q", test.testcase, test.input, test.expected, actual)
			t.Fail()
		}
	}

}

func areEqual(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a[0]) != len(b[0]) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
