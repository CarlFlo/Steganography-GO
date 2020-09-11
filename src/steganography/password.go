package steganography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

/*
	Encrypt will encrypt the provided data with the password

	input:
		password string : The password to encrypt with
		data *[]byte : The data to be encrypted

	output:
		error : Error if any
*/
func Encrypt(password string, data *[]byte) error {

	key := []byte(password)

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return err
	}

	*data = gcm.Seal(nonce, nonce, *data, nil)

	return nil
}

/*
	Decrypt will decrypt the provided data with the password

	input:
		password string : The password to decrypt with
		data *[]byte : The data to be decrypted

	output:
		error : Error if any
*/
func Decrypt(password string, data *[]byte) error {

	key := []byte(password)

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return err
	}
	nonce, ciphertext := (*data)[:gcm.NonceSize()], (*data)[gcm.NonceSize():]
	*data, err = gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	return nil
}

/*
	KeySanitizing will sanitize the key to make sure it's
	not too long or short. Very basic

	input:
		key string : The key to be used in encryption/decryption

	output:
		string : The valid key
*/
func KeySanitizing(key string) string {

	if len(key) < 16 {
		for len(key) < 16 {
			key += "-"
		}

	} else if len(key) > 32 {
		key = key[:32]
	}

	return key
}
