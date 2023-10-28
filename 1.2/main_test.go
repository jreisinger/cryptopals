package main

import "testing"

func TestFixedXOR(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	seed := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"
	got, err := FixedXOR(input, seed)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
