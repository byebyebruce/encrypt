package encrypt

const (
	mulKey uint16 = 51423
)

// Encrypt 加密
func Encrypt(buffer []byte, key uint16) {
	for i := 0; i < len(buffer); i += 2 {
		buffer[i] = buffer[i] ^ byte(key>>8)
		key = (uint16(int8(buffer[i])) + key) * mulKey
	}
}

// Decrypt 解密
func Decrypt(buffer []byte, key uint16) {
	for i := 0; i < len(buffer); i += 2 {
		c := buffer[i]
		buffer[i] = c ^ byte(key>>8)
		key = (uint16(int8(c)) + key) * mulKey
	}
}
