package main

import (
	"fmt"
	"unicode"
)

type plaintext struct {
	key         byte
	plaintext   string
	englishness int
}

type byEnglishness []plaintext

func (x byEnglishness) Len() int           { return len(x) }
func (x byEnglishness) Less(i, j int) bool { return x[i].englishness < x[j].englishness }
func (x byEnglishness) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func printPlaintexts(plaintexts []plaintext) {
	for _, r := range plaintexts {
		fmt.Printf("(%d) %q -> %q\n", r.englishness, r.key, r.plaintext)
	}
}

// bruteforce gets plaintexts from ciphertext by xoring it byte by byte against
// all 256 ASCII characters. It also calculates the englishness of the obtained
// plaintext.
func bruteforce(ciphertext []byte) []plaintext {
	var plaintexts []plaintext
	for i := 0; i < 256; i++ { // all ASCII characters
		key := byte(i)
		s := xor(ciphertext, key)
		n := englishness(s)
		plaintexts = append(plaintexts, plaintext{
			key:         key,
			plaintext:   s,
			englishness: n,
		})
	}
	return plaintexts
}

// englishness counts how many of the most frequent english characters are
// present in s.
func englishness(s string) int {
	var count int
	for _, c := range s {
		switch unicode.ToUpper(c) {
		case 'E', 'T', 'A', 'O', 'I', 'N', ' ', 'S', 'H', 'R', 'D', 'L', 'U':
			count++
		}
	}
	return count
}

// xor xors all bytes of input against key.
func xor(input []byte, key byte) string {
	out := make([]byte, len(input))
	for i := range input {
		out[i] = input[i] ^ key
	}
	return string(out)
}
