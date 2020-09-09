package steganography

import (
	"fmt"
	"strconv"
)

/*
	binaryStringToByteArray converts a string to a []byte containing its binary data

	input:
		s string : A string with binary data,
			T.ex. "01110100011001010111001101110100" ('test' in binary)

	output:
		[]byte : returns the binary data as 1's and 0's
*/
func binaryStringToByteArray(s string) []byte {

	// Not very efficient. Uses 8 times more space than needed

	data := make([]byte, len(s))

	for i, c := range s {
		if c == 48 { // 0
			data[i] = 0
		} else { // 1
			data[i] = 1
		}
	}

	return data
}

/*
	stringToBinary converts a string to binary
	input:
		str string : A string to be converted to binary
			T.ex. Input "test"

	output:
		byte[] : Returns the string as a []byte slice containing the binary data
			T.ex. Output [0,1,1,1,0,1,0,0,0,1,1,0,0,1,0,1,0,1,1,1,0,0,1,1,0,1,1,1,0,1,0,0]
*/
func stringToBinary(str string) []byte {

	buffer := make([]byte, (len(str) * 8))

	for index, c := range str {
		for i := 0; i < 8; i++ {
			bit := c & (0x80 >> uint(i))

			if bit > 0 {
				bit = 1
			}

			buffer[index*8+i] = byte(bit)
		}
	}

	return buffer
}

/*
	intToUint32Binary convers a int to a uint32 value first
	and then returns a []byte slice with value in binary form.

	input:
		number int : The number to be converted into a binary []byte slice

	output:
		[]byte : The binary byte slice
*/
func intToUint32Binary(number int) []byte {

	n := int64(number)
	asString := fmt.Sprintf("%032b", n)
	asByte := binaryStringToByteArray(asString)

	return asByte
}

/*
	binaryStringToInt takes a string containing binary data and converts it
	to an integer

	input:
		data string : The string with the binary data

	output:
		int : The int value
		error : The error if any
*/
func binaryStringToInt(data string) (int, error) {
	value, err := strconv.ParseInt(data, 2, 64)
	if err != nil {
		return -1, err
	}

	return int(value), nil
}
