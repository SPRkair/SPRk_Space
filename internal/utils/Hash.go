package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

type HashType string

const Sha256 HashType = "sha256"

func Hash(hashType HashType, line string) string {
	switch hashType {
	case Sha256:
		sum := sha256.Sum256([]byte(line))
		tokenHash := hex.EncodeToString(sum[:])
		return tokenHash
	default:
		return line
	}
}
