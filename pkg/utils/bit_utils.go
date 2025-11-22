package utils

// StringToBits converts string to bit slice (MSB first)
func StringToBits(s string) []bool {
	bits := make([]bool, 0, len(s)*8)
	for _, r := range s {
		byteVal := byte(r)
		for i := 7; i >= 0; i-- {
			bits = append(bits, (byteVal>>i)&1 == 1)
		}
	}
	return bits
}

// BitsToString converts bit slice to string, trimming to exact byte count
func BitsToString(bits []bool, charCount int) string {
	// Обрезаем до нужного количества байт (charCount * 8)
	trimLen := charCount * 8
	if len(bits) < trimLen {
		trimLen = len(bits) - (len(bits) % 8) // Берём полные байты
	}
	bits = bits[:trimLen]

	var result []byte
	for i := 0; i < len(bits); i += 8 {
		var b byte
		for j := 0; j < 8; j++ {
			if bits[i+j] {
				b |= (1 << (7 - j))
			}
		}
		result = append(result, b)
	}
	return string(result)
}
