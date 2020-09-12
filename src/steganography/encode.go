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
	Encode will take data to be encrypted into the provided image

	input:
		data []byte: The binary []byte data to be encoded
		img image.Image : The image to have the message encoded inside it
		outFile string : The name of the file that will be created with the encoded message

	output:
		error : If there was an error
*/
func Encode(data []byte, img image.Image, outFile, password string) error {

	// Encrypts the message
	if err := Encrypt(password, &data); err != nil {
		return err
	}

	// Concats 4 bytes to the message containing the length of the message
	addDataLengthToData(&data)

	if err := checkAvaiableSize(&data, img); err != nil {
		return err
	}

	newImg := makeImageEditable(img)

	rowL := newImg.Bounds().Max.Y
	colL := newImg.Bounds().Max.X

	// Content in data is byte and not binary data
	// Rework this function

	ch := make(chan bool)

	go func(ch chan bool, data *[]byte) {
		for _, _byte := range *data {
			for i := 0; i < 8; i++ {
				result := _byte & (0x80 >> i)
				if result > 0 {
					ch <- true
				} else {
					ch <- false
				}
			}
		}
		close(ch)
	}(ch, &data)

	i := 0
	for bitSet := range ch {
		row := (i / colL) % rowL
		col := i % colL

		// Gets the pixel
		rgbaArray := getRGBAArray(newImg.At(row, col))

		/* */

		newImg.Set(row, col, color.RGBA{rgbaArray[0], rgbaArray[1], rgbaArray[2], rgbaArray[3]})
		i++
	}

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
