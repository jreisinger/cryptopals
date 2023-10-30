package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"sort"
)

func main() {
	ciphertext, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		fmt.Fprintf(os.Stderr, "1.3: %v\n", err)
		os.Exit(1)
	}
	plaintexts := bruteforce(ciphertext)
	sort.Sort(byEnglishness(plaintexts))
	printPlaintexts(plaintexts)
}
