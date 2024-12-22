package blocksize

func padWithZeros(block []byte, desiredSize int) []byte {
	compleateBlock := make([]byte, desiredSize)

	for i, b := range block {
		compleateBlock[i] = b
	}

	for i := len(block); i < desiredSize; i++ {
		compleateBlock[i] = 0x00
	}

	return compleateBlock
}
