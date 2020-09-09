package steganography

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

/*
	LoadImage loads the image with the provided filename and returns
	a image object and error (if any). Can handle multible fileformats

	input:
		fName string : The filename string

	output:
		image.Image : The image data
		error : If there was an error

*/
func LoadImage(fName string) (image.Image, error) {

	file, err := os.Open(fmt.Sprintf("./%s", fName))
	if err != nil {
		return nil, err
	}

	// Gets the files extension from the provided filename (removed the dot as well)
	fileExtension := strings.ToLower(filepath.Ext(fName)[1:])
	// Decodes the image
	switch fileExtension {
	case "png":
		return png.Decode(file)
	case "jpg":
		fallthrough
	case "jpeg":
		return jpeg.Decode(file)
	case "gif":
		return gif.Decode(file)
	default:
		return nil, fmt.Errorf("%s is unsupported", fileExtension)
	}

}
