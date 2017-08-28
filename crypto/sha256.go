package crypto

import(
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(s string) string{
	byteArray := sha256.Sum256([]byte(s))
	return hex.EncodeToString(byteArray[:]);
}