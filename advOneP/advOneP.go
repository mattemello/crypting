package advonep

func crypt(textCh, keyCh <-chan byte, result chan<- byte) {
	defer close(result)
	for {
		bText, more := <-textCh
		bKey, moreK := <-keyCh
		if moreK && more {
			result <- (bText ^ bKey)
		} else {
			return
		}
	}
}
