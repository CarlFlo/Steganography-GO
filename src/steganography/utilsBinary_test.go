package steganography

// https://www.youtube.com/watch?v=sOeUf1YICSA

import (
	"bytes"
	"testing"

	"../myTest"
)

func TestBinaryStringToByteArray(t *testing.T) {

	tvh := myTest.TestValueHolder{}
	tvh.AddTestValue("01110100011001010111001101110100", []byte{0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0})

	for _, e := range tvh.Tests {
		input := e.Input.(string)
		expectedOutput := e.ExpectedOutput.([]byte)

		actual := binaryStringToByteArray(input)

		res := bytes.Compare(actual, expectedOutput)
		if res != 0 {
			t.Errorf(e.FormatError(actual))
		}
	}

}

func TestStringToBinary(t *testing.T) {

	tvh := myTest.TestValueHolder{}
	tvh.AddTestValue("test", []byte{0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 1, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 0})

	for _, e := range tvh.Tests {

		input := e.Input.(string)
		expectedOutput := e.ExpectedOutput.([]byte)

		actual := stringToBinary(input)
		res := bytes.Compare(actual, expectedOutput)
		if res != 0 {
			t.Errorf(e.FormatError(actual))
		}
	}

}

func TestIntToUint32Binary(t *testing.T) {

	tvh := myTest.TestValueHolder{}
	tvh.AddTestValue(5324, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0})
	tvh.AddTestValue(0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})

	for _, e := range tvh.Tests {
		input := e.Input.(int)
		expectedOutput := e.ExpectedOutput.([]byte)

		actual := intToUint32Binary(input)
		res := bytes.Compare(actual, expectedOutput)
		if res != 0 {
			t.Errorf(e.FormatError(actual))
		}
	}

}

func TestBinaryStringToInt(t *testing.T) {
	tvh := myTest.TestValueHolder{}
	tvh.AddTestValue("00000000", 0)
	tvh.AddTestValue("00001010", 10)
	tvh.AddTestValue("11111111", 255)

	for _, e := range tvh.Tests {
		input := e.Input.(string)
		expectedOutput := e.ExpectedOutput.(int)

		actual, err := binaryStringToInt(input)
		if err != nil {
			t.Errorf(err.Error())
		}

		if expectedOutput != actual {
			t.Errorf(e.FormatError(actual))
		}
	}
}
