package steganography

/*
	addDataLengthToData adds the length as a binary []byte slice to the message data

	input:
		data *[]byte: Pointer to the data

*/
func addDataLengthToData(data *[]byte) {

	//length := intToUint32Binary(len(*data))

	byteLength := make([]byte, 4)
	length := uint32(len(*data))

	for i := 0; i < len(byteLength); i++ {
		for j := 0; j < 8; j++ {
			shift := 8*i + j
			res := length & (0x80000000 >> shift)
			if res > 0 {
				res = 1
			}

			byteLength[i] |= byte(res) << (7 - j)
		}
	}

	*data = append(byteLength, *data...)
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
