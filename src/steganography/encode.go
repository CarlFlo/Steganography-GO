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

		rgbaArray := getRGBAArray(newImg.At(row, col))

		for j := i; j < i+3 && j < len(data); j++ {
			setLSB(&rgbaArray[j-i], data[j]&1 == 1)
		}

		// Inserts the new color into the pixel
		newImg.Set(row, col, color.RGBA{rgbaArray[0], rgbaArray[1], rgbaArray[2], rgbaArray[3]})
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
