package onetimepad

func crypt(plaintext, key []byte) []byte {
	var crypted = make([]byte, len(key))

	for i, _ := range plaintext {
		crypted[i] = plaintext[i] ^ key[i]
	}

	return crypted
}
