package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

// 签名验证
func VerifySignature(text string, signature string, secret string) bool {
	h := hmac.New(sha1.New, []byte(secret))
	h.Write([]byte(text))
	textSignature := "sha1=" + hex.EncodeToString(h.Sum(nil))
	return signature == textSignature
}
