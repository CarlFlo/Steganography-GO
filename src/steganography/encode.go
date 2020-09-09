package steganography

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
)

/*
	EncodeString is a wrapper function for Encode, that handles strings

	input:
		message string: The binary []byte data to be encoded
		img image.Image : The image to have the message encoded inside it
		outFile string : The name of the file that will be created with the encoded message

	output:
		error : If there was an error
*/
func EncodeString(message string, img image.Image, outFile string) error {
	data := stringToBinary(message)
	return Encode(data, img, outFile)
}

/*
	Encode will take data to be encrypted into the provided image

	input:
		data []byte: The binary []byte data to be encoded
		img image.Image : The image to have the message encoded inside it
		outFile string : The name of the file that will be created with the encoded message

	output:
		error : If there was an error
*/
func Encode(data []byte, img image.Image, outFile string) error {

	addDataLengthToData(&data)
	// The first 32 bits of the binary data contains the length
	// The rest is the binary data

	if err := checkAvaiableSize(&data, img); err != nil {
		return err
	}

	newImg := makeImageEditable(img)

	rowL := newImg.Bounds().Max.Y
	colL := newImg.Bounds().Max.X

	// Iterate over rows and cols, increment by 3 because each pixel can hold 3 "values"
	for i := 0; i < len(data); i += 3 {
		row := (i / colL) % rowL
		col := i % colL

		r, g, b, a := getRGBA(newImg.At(row, col))

		rgbaArray := []uint8{r, g, b, a}

		for j := i; j < i+3 && j < len(data); j++ {
			setLSB(&rgbaArray[j-i], data[j]&1 == 1)
		}
		/* // Old implementation
		// Check if the least significant bit needs to be changed
		setLSB(&r, data[i]&1 == 1)
		if i+1 < len(data) { // Check for out of bounds. We do 3 increments at a time so we have to
			setLSB(&g, data[i+1]&1 == 1)
		}
		if i+2 < len(data) { // Check for out of bounds
			setLSB(&b, data[i+2]&1 == 1)
		}
		*/
		// Inserts the new color into the pixel
		newImg.Set(row, col, color.RGBA{rgbaArray[0], rgbaArray[1], rgbaArray[2], rgbaArray[3]})
		//newImg.Set(row, col, color.RGBA{r, g, b, a})
	}

	// Removes extension
	outFile = outFile[:len(outFile)-len(filepath.Ext(outFile))]

	// Write to file
	if outFile, err := os.Create(fmt.Sprintf("%s_changed.png", outFile)); err != nil {
		return err
	} else {
		png.Encode(outFile, newImg)
		outFile.Close()
	}

	return nil
}
