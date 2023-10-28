package main

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(s string) (string, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(b), nil
}
