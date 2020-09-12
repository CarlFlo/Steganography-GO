package steganography

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

/*
	makeImageEditable makes a copy of the provided image that can be edited

	input:
		img image.Image : The image to be made editable

	output:
		*image.RGBA : A copy of the input image that is editable
*/
func makeImageEditable(img image.Image) *image.RGBA {

	// Makes image drawable
	newImg := image.NewRGBA(img.Bounds())
	draw.Draw(newImg, img.Bounds(), img, image.Point{}, draw.Over)

	return newImg
}

/*
	getRGBA converts a color.Color object to its rgba uint8 values

	input:
		c color.Color : The color to be converted

	output:
		The four RGB uint8 values
*/
func getRGBA(c color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := c.RGBA()

	constant := uint32(65535 / 255) // 65535 / 255 = 257

	newR := uint8(r / constant)
	newG := uint8(g / constant)
	newB := uint8(b / constant)
	newA := uint8(a / constant)

	return newR, newG, newB, newA
}

/*
	Wrapper function for getRGBA that returns the values in an array

	input:
		c color.Color : The color to be converted

	output:
		The four RGB uint8 values in an uint8 array
*/
func getRGBAArray(c color.Color) []uint8 {

	r, g, b, a := getRGBA(c)
	return []uint8{r, g, b, a}
}

/*
	checkAvaiableSize checks so that the provided image can contain the message

	input:
		message *[]byte : The points to the message to be encoded
		img image.Image : The image that will have the message encoded in it

	output:
		error : The error if the image size isn't enough
*/
func checkAvaiableSize(message *[]byte, img image.Image) error {

	// Multiply by 3 because each pixel has 3 values (RGB)
	// Where 1 bit can be encoded in their respective 3 RGB values in the LSB
	maxSize := img.Bounds().Max.Y * img.Bounds().Max.X * 3 * 1

	neededSize := len(*message) * 8

	// Check if the image is big enough for the message
	if neededSize > maxSize {
		errMsg := fmt.Sprintf("The image is not big enough to encrypt the message. Avaiable: %v, Needed: %v, diff: %v", maxSize, neededSize, neededSize-maxSize)
		return errors.New(errMsg)
	}
	return nil
}
