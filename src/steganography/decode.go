package steganography

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"strconv"
	"strings"
)

/*
	Decode decodes the hidden message inside the image

	input:
		img image.Image : The image with the secret message
*/
func Decode(img image.Image, password string) (string, error) {

	// Holds the message
	var buffer bytes.Buffer

	rowL := img.Bounds().Max.Y
	colL := img.Bounds().Max.X
	maxSize := rowL * colL
	iterateTo := maxSize
	hasLength := false

	// Iterate over rows and cols, increment by 3 because each pixel can hold 3 "values"
	for i := 0; i < iterateTo; i += 3 {
		row := (i / colL) % rowL
		col := i % colL

		r, g, b, _ := getRGBA(img.At(row, col))

		// Saves the binary (1 or 0) data in the buffer
		buffer.WriteString(fmt.Sprintf("%v%v%v", r&1, g&1, b&1))

		// Checks if we have found the size of the message
		if !hasLength && len(buffer.Bytes()) >= 32 {
			hasLength = true

			// Take the symbold 0-31 (32) and converts them to an int
			if newSize, err := binaryStringToInt(buffer.String()[:32]); err != nil {
				return "", err
			} else {

				// Update how long we need to iterate
				iterateTo = newSize + 32

				fmt.Println("Found length", iterateTo-32)

				// This check checks if the image could have something encrypted in it
				// If 'newSize' is larger than maxSize then we throw an error
				// because no message of that size could be encoded in that image
				// Multiply by 3 because each pixel can 'hold' 3 bits
				if iterateTo > maxSize*3 {
					return "", errors.New("The provided image is invalid")
				}
			}
		}
	}

	// Takes the result, removes all access information by 'substringing' it away.
	// The buffer holds the values 48(0) and 49(1). binaryStringToByteArray converts the values to the numbers
	decodedResult := binaryStringToByteArray(buffer.String()[32:iterateTo])

	/*
		// Decrypt with the password
		if err := Decrypt(password, &asd); err != nil {
			return "", err
		}

	*/
	// This converts the []byte with the binary data to a string with the message
	result := toString(prettyPrint(decodedResult))

	return result, nil
}

/*
	toString converts a binarystring to a textstring

	input:
		data string : The string containing binary data

	output:
		string : A string containing binary 0&1 as a readable string
*/
func toString(data string) string {
	buffer := make([]byte, 0)
	for _, s := range strings.Fields(data) {
		n, _ := strconv.ParseUint(s, 2, 8)
		buffer = append(buffer, byte(n))
	}
	return string(buffer)
}
