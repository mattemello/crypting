package keyschedules

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	var masterRoundedKey [4]byte
	for i, b := range masterKey {
		masterRoundedKey[i] = (b ^ byte(roundNumber))
	}

	return masterRoundedKey
}
