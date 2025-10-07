package utils

import "strconv"

func UintToTextBytes(n uint) []byte {
	return []byte(strconv.FormatUint(uint64(n), 10))
}
