package sbox

import "errors"

func sBox(b byte) (byte, error) {
	workByts := b
	lastHalf := 0b11 & workByts
	workByts >>= 2
	firstHalf := 0b11 & workByts
	workByts >>= 2

	if firstHalf == lastHalf {
		return 0b00, nil
	} else if (firstHalf == 0b01 && lastHalf == 0b00) || (lastHalf == 0b01 && firstHalf == 0b00) {
		return 0b10, nil
	} else if (firstHalf == 0b10 && lastHalf == 0b00) || (lastHalf == 0b10 && firstHalf == 0b00) {
		return 0b01, nil
	} else if (firstHalf == 0b11 && lastHalf == 0b00) || (lastHalf == 0b11 && firstHalf == 0b00) {
		return 0b11, nil
	} else if (firstHalf == 0b01 && lastHalf == 0b10) || (lastHalf == 0b01 && firstHalf == 0b10) {
		return 0b11, nil
	} else if (firstHalf == 0b01 && lastHalf == 0b11) || (lastHalf == 0b01 && firstHalf == 0b11) {
		return 0b11, nil
	} else if (firstHalf == 0b10 && lastHalf == 0b11) || (lastHalf == 0b10 && firstHalf == 0b11) {
		return 0b10, nil
	} else {
		return 0b00, errors.New("invalid input")
	}
}
