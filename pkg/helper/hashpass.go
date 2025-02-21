package helper

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// GenerateFromPassword menghasilkan hash dari password menggunakan Argon2id.
// Format hash yang dihasilkan: $argon2id$v=19$t=<time>$m=<memory>$p=<threads>$<salt>$<hash>
func ArgonGeneratePassword(password string) (string, error) {
	const saltLength = 16
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Parameter Argon2id
	const timeParam = 1           // iterasi
	const memoryParam = 64 * 1024 // memori dalam KB (misal: 64 MB)
	const threadsParam = 4        // jumlah thread (parallelism)
	const keyLen = 32             // panjang hash yang dihasilkan

	// Generate hash
	hash := argon2.IDKey([]byte(password), salt, timeParam, memoryParam, uint8(threadsParam), keyLen)

	// Encode salt dan hash ke base64 tanpa padding
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Buat string encoded yang menyimpan parameter dan hasil hash
	encodedHash := fmt.Sprintf("$argon2id$v=19$t=%d$m=%d$p=%d$%s$%s", timeParam, memoryParam, threadsParam, b64Salt, b64Hash)

	return encodedHash, nil
}

func ArgonComparePassword(encodedHash, password string) (bool, error) {
	// Split string encoded dengan separator "$"
	// Format yang diharapkan: ["", "argon2id", "v=19", "t=1", "m=65536", "p=4", "<salt>", "<hash>"]
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 8 {
		return false, fmt.Errorf("format hash tidak valid")
	}

	// Parse parameter
	var timeParam uint32
	var memoryParam uint32
	var threadsParam uint8

	_, err := fmt.Sscanf(parts[3], "t=%d", &timeParam)
	if err != nil {
		return false, err
	}
	_, err = fmt.Sscanf(parts[4], "m=%d", &memoryParam)
	if err != nil {
		return false, err
	}
	_, err = fmt.Sscanf(parts[5], "p=%d", &threadsParam)
	if err != nil {
		return false, err
	}

	// Decode salt dan hash yang tersimpan
	salt, err := base64.RawStdEncoding.DecodeString(parts[6])
	if err != nil {
		return false, err
	}
	hashStored, err := base64.RawStdEncoding.DecodeString(parts[7])
	if err != nil {
		return false, err
	}

	// Hitung hash dari password input dengan parameter yang sama
	hashComputed := argon2.IDKey([]byte(password), salt, timeParam, memoryParam, threadsParam, uint32(len(hashStored)))

	// Gunakan perbandingan konstan waktu untuk menghindari timing attack
	if subtle.ConstantTimeCompare(hashStored, hashComputed) == 1 {
		return true, nil
	}
	return false, nil
}
