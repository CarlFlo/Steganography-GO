package steganography

import (
	"bytes"
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

	/*
		if err := Encrypt(password, &data); err != nil {
			return err
		}
	*/

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

	var buffer bytes.Buffer

	/* Message encoding is broken, check how the lsb is set */

	/* The encoded data is the same compared to the 'old' version,
	but the method of encoding is not working */

	/* Encoding is down each column and then it goes down to the bottom. I know weird from code setup*/

	exitFlag := false
	for row := 0; row < rowL; row++ {
		for col := 0; col < colL; col++ {

			rgbaArray := getRGBAArray(newImg.At(row, col))

			for i := 0; i < 3; i++ {
				/* Apply change here */

				/*
					if bitSet, ok := <-ch; !ok {
						exitFlag = true
						break
					} else {
						setLSB(&rgbaArray[i], bitSet)
						buffer.WriteString(fmt.Sprintf("%v", rgbaArray[i]&1))
					}
				*/

				bitSet, ok := <-ch
				if !ok {
					exitFlag = true
					break
				}
				setLSB(&rgbaArray[i], bitSet)
				buffer.WriteString(fmt.Sprintf("%v", rgbaArray[i]&1))
			}

			newImg.Set(row, col, color.RGBA{rgbaArray[0], rgbaArray[1], rgbaArray[2], rgbaArray[3]})
			if exitFlag {
				goto Finished
			}
		}
	}

	// No more data to encode in the image
Finished:

	fmt.Println(buffer.String()[:32], "-", buffer.String()[32:])

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
