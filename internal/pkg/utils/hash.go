package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash256(words ...any) string {
	payload := fmt.Sprint(words...)
	h := sha256.New()
	h.Write([]byte(payload))
	return hex.EncodeToString(h.Sum(nil))
}
