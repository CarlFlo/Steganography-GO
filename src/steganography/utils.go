package steganography

import (
	"bytes"
	"fmt"
)

/*
	prettyPrint will format the provided []byte as a string that is "pretty"

	input:
		data []byte : The byte slice data to be formatted

	output:
		string : The formatted string
*/
func prettyPrint(data []byte) string {

	var buffer bytes.Buffer

	for i := 0; i < len(data); i += 8 {
		for j := i; j < i+8; j++ {
			buffer.WriteString(fmt.Sprintf("%d", data[j]))
		}
		buffer.WriteString(" ")
	}
	return buffer.String()
}
