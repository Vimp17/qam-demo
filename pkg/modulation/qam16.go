package modulation

type Complex struct {
	I, Q float64
}

// Gray-coded QAM-16 mapping (bits: [b0, b1, b2, b3])
func Modulate16(bits []bool) []Complex {
	// Дополняем до кратности 4 битам
	for len(bits)%4 != 0 {
		bits = append(bits, false)
	}

	symbols := make([]Complex, 0, len(bits)/4)
	for i := 0; i < len(bits); i += 4 {
		b0, b1, b2, b3 := bits[i], bits[i+1], bits[i+2], bits[i+3]

		// Gray mapping для I (b0,b1):
		// 00 -> +3, 01 -> +1, 11 -> -1, 10 -> -3
		var I float64
		if b0 == b1 { // 00 или 11
			if !b0 {
				I = 3.0 // 00
			} else {
				I = -1.0 // 11
			}
		} else { // 01 или 10
			if !b0 {
				I = 1.0 // 01
			} else {
				I = -3.0 // 10
			}
		}

		// Gray mapping для Q (b2,b3) - аналогично
		var Q float64
		if b2 == b3 {
			if !b2 {
				Q = 3.0
			} else {
				Q = -1.0
			}
		} else {
			if !b2 {
				Q = 1.0
			} else {
				Q = -3.0
			}
		}

		symbols = append(symbols, Complex{I: I, Q: Q})
	}
	return symbols
}

func Demodulate16(symbols []Complex) []bool {
	bits := make([]bool, 0, len(symbols)*4)
	for _, s := range symbols {
		// Решение для I
		var b0, b1 bool
		if s.I > 0 {
			if s.I > 2 {
				b0, b1 = false, false // +3
			} else {
				b0, b1 = false, true // +1
			}
		} else {
			if s.I < -2 {
				b0, b1 = true, false // -3
			} else {
				b0, b1 = true, true // -1
			}
		}

		// Решение для Q (аналогично)
		var b2, b3 bool
		if s.Q > 0 {
			if s.Q > 2 {
				b2, b3 = false, false
			} else {
				b2, b3 = false, true
			}
		} else {
			if s.Q < -2 {
				b2, b3 = true, false
			} else {
				b2, b3 = true, true
			}
		}

		bits = append(bits, b0, b1, b2, b3)
	}
	return bits
}
