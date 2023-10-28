package main

import "encoding/hex"

func FixedXOR(hexa, seed string) (string, error) {
	b, err := hex.DecodeString(hexa)
	if err != nil {
		return "", err
	}
	s, err := hex.DecodeString(seed)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(b))
	for i := range b {
		out[i] = b[i] ^ s[i]
	}
	return hex.EncodeToString(out), nil
}
