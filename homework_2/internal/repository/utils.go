package repository

import "math/rand"

func generateId() string {
	symbols := make([]byte, 10)
	for i := range symbols {
		switch rand.Intn(3) {
		case 0:
			symbols[i] = byte(rand.Intn(26)) + 'a' // 26 букв от 'a' до 'z'
		case 1:
			symbols[i] = byte(rand.Intn(26)) + 'A' // 26 букв от 'A' до 'Z'
		case 2:
			symbols[i] = byte(rand.Intn(10)) + '0' // 10 цифр от '0' до '9'
		}
	}
	return string(symbols)
}
