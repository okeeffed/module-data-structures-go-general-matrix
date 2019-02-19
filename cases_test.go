package caps

type testCase struct {
	description string
	input       int
	expected    [][]int
}

var testCases = []testCase{
	{
		description: "should generate 1x1 matrix",
		input:       1,
		expected:    [][]int{{0}},
	},
	{
		description: "should generate 2x2 matrix",
		input:       2,
		expected:    [][]int{{0, 1}, {2, 3}},
	},
	{
		description: "should return 3x3 matrix",
		input:       3,
		expected:    [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
	},
}
