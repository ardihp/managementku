package helpers

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateUUIDv4() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Atur bit versi ke 4 (0100)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// Atur bit variant ke RFC 4122 (10xx)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16],
	), nil
}

func GenerateUUIDv7() (string, error) {
	var uuid [16]byte

	// 1. Dapatkan timestamp saat ini dalam milidetik (48 bit pertama)
	milli := time.Now().UnixMilli()

	uuid[0] = byte(milli >> 40)
	uuid[1] = byte(milli >> 32)
	uuid[2] = byte(milli >> 24)
	uuid[3] = byte(milli >> 16)
	uuid[4] = byte(milli >> 8)
	uuid[5] = byte(milli)

	// 2. Isi sisa 10 byte (byte 6 sampai 15) dengan nilai acak kriptografis
	_, err := rand.Read(uuid[6:])
	if err != nil {
		return "", fmt.Errorf("gagal membaca random bytes: %w", err)
	}

	// 3. Set Version ke 7 (0111) pada 4-bit teratas byte ke-6
	uuid[6] = (uuid[6] & 0x0f) | 0x70

	// 4. Set Variant ke RFC 9562/4122 (10xx) pada 2-bit teratas byte ke-8
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format menjadi representasi string canonical (8-4-4-4-12)
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4],
		uuid[4:6],
		uuid[6:8],
		uuid[8:10],
		uuid[10:16],
	), nil
}
