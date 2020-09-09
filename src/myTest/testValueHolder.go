package myTest

import "fmt"

// TestValue holds each test. The input and expected output with the ID of the test
type TestValue struct {
	ID             uint8
	Input          interface{} // Simple. Only 1 input for now
	ExpectedOutput interface{}
}

// FormatError formats the error message
func (TV *TestValue) FormatError(actual interface{}) string {
	return fmt.Sprintf("(ID: %d) expected '%v', but got '%v'. Input was '%v'", TV.ID, TV.ExpectedOutput, actual, TV.Input)
}

// TestValueHolder keeps track of the amount of
type TestValueHolder struct {
	Tests []TestValue
}

// AddTestValue adds the input and expected output to the struct
func (TVH *TestValueHolder) AddTestValue(input interface{}, expectedOutput interface{}) {
	tv := TestValue{ID: uint8(len(TVH.Tests)), Input: input, ExpectedOutput: expectedOutput}
	TVH.Tests = append(TVH.Tests, tv)
}

/* Usage

tvh := myTest.TestValueHolder{}
tvh.AddTestValue("The input", "Expected output")

for _, e := range tvh.Tests {
	input := e.Input.(DESIRED_TYPE_HERE)
	expectedOutput := e.ExpectedOutput.(DESIRED_TYPE_HERE)

	actual := THE_FUNCTION(INPUT)

	res := THE_COMPARE_FUNCTION(actual, expectedOutput)
	if res != 0 { // If they don't match
		t.Errorf(e.FormatError(actual))
	}
}
*/
