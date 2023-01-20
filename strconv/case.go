package strconv

func ByteLower(c byte) byte {
	return c | ('x' - 'X')
}

func ByteUpper(c byte) byte {
	if c < 'a' || c > 'z' {
		return c
	}
	const b uint8 = 32
	return c & (^b)
}
