package utils

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func EncodeToBase62(num int) string {
	if num == 0 {
		return "000000"
	}

	encoded := ""
	base := len(base62Chars)

	for num > 0 {
		remainder := num % base
		encoded = string(base62Chars[remainder]) + encoded
		num = num / base
	}

	// Pad to 6 characters
	for len(encoded) < 6 {
		encoded = "0" + encoded
	}

	return encoded
}

func DecodeFromBase62(str string) int {
	decoded := 0
	base := len(base62Chars)

	for _, char := range str {
		decoded = decoded * base
		for i, c := range base62Chars {
			if c == char {
				decoded += i
				break
			}
		}
	}

	return decoded
}