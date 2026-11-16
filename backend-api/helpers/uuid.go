package helpers

import (
	"crypto/rand"
	"fmt"
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
