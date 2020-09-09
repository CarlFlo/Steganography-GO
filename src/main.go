package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"./steganography"
)

var flagEncrypt bool
var flagDecrypt bool
var flagMessage string
var flagFilepath string

func init() {
	flag.BoolVar(&flagEncrypt, "e", false, "Encrypt flag. Set if you want to encrypt")
	flag.BoolVar(&flagDecrypt, "d", false, "Decrypt flag. Set if you want to decrypt")
	flag.StringVar(&flagMessage, "m", "", "Message: the text to be encrypted here surrounded by quotes")
	flag.StringVar(&flagFilepath, "f", "", "File: the path to the image (jpg, jpeg, png & gif is supported)") // jpg == jpeg

	flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "e":
			flagEncrypt = true
		case "d":
			flagDecrypt = true
		}
	})

	flag.Parse()
}

func main() {

	if err := run(); err != nil {
		panic(err.Error())
	}

}

/*
	run executes the program

	output:
		error : Error if there is an error
*/
func run() error {
	// Needs to be run from commandline

	// Error checking
	if len(os.Args) == 1 {
		// Help message
		fmt.Println("You cannot run this program from without providing any arguments")
		fmt.Println(fmt.Sprintf("Run: '%s -h' from the commandline to get help", filepath.Base(os.Args[0])))
		fmt.Println("Press the Enter Key to continue...")
		fmt.Scanln() // wait for Enter Key
		return errors.New("You cannot launch the program without providing any arguments")
	} else if flagEncrypt && flagDecrypt {
		return errors.New("Both encrypt and decrypt flag cannot be set")
	} else if !flagEncrypt && !flagDecrypt {
		return errors.New("Both encrypt and decrypt flag aren't set")
	} else if len(flagFilepath) == 0 {
		return errors.New("A filepath needs to be provided")
	} else if flagEncrypt && len(flagMessage) == 0 {
		return errors.New("You have chosen to encrypt but havent provided a message to be encrypted")
	}

	var result string
	var err error

	if flagEncrypt { // If true then encrypt
		result, err = encrypt(flagFilepath, flagMessage)
	} else if flagDecrypt { // If true then decrypt
		result, err = decrypt(flagFilepath)
	}

	if err == nil {
		fmt.Println(result)
	}

	return err
}

/*
	encrypt will hide the provided message inside the image

	input:
		fName string : The filename of the image
		message string : The message to be hidden in the image

	output:
		string : The output
		error : Error if there is an error
*/
func encrypt(fName, message string) (string, error) {

	// Loads png image
	data, err := steganography.LoadImage(fmt.Sprintf("%s", fName))
	if err != nil {
		return "", err
	}

	// Performs steganography
	if err = steganography.EncodeString(message, data, fmt.Sprintf("%s_changed", fName)); err != nil {
		return "", err
	}

	return fmt.Sprintf("Sucessfully encoded the message: '%s'", message), nil
}

/*
	decrypt will extract the hidden message inside a png image

	input:
		fName string : The filename of the png image

	output:
		string : The extracted message
		error : Error if there is an error
*/
func decrypt(fName string) (string, error) {

	// Loads png image
	data, err := steganography.LoadImage(fmt.Sprintf("%s", fName))
	if err != nil {
		return "", err
	}

	// Performs steganography
	result, err := steganography.Decode(data)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("The decoded message is: '%s'", result), nil
}

// Test preforms a test
// This is true
func Test() error {

	return nil
}
