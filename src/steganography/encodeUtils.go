package steganography

/*
	addDataLengthToData adds the length as a binary []byte slice to the message data

	input:
		data *[]byte: Pointer to the data

*/
func addDataLengthToData(data *[]byte) {

	length := intToUint32Binary(len(*data))
	*data = append(length, *data...)
}

/*
	setLSB sets the least significant bit

	input:
		n *uint8 : Pointer to the value whos bit will be set to match with the provided value
		set bool : If the LSB should be set or not
*/
func setLSB(n *uint8, set bool) {

	// Checks if the LSB bit is set for the number provided
	isSet := (*n&1 == 1)

	// Toggle the LSB if they dont match
	if isSet != set {
		*n ^= 1 // Toggle
	}
}
