package basic

import (
	"cryptopals/assert"
	"fmt"
	"os"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	got, err := HexToBase64(input)
	assert.Nil(t, err)
	assert.Equal[string](t, got, want)
}

func TestFixedXOR(t *testing.T) {
	t.Run("length not equal", func(t *testing.T) {
		want, err := FixedXOR("1c0111001f0101", "1c0111001f01011112")
		assert.Equal[error](t, err, ErrLengthNotEqual)
		assert.True(t, len(want) == 0)
	})

	t.Run("XORs correctly", func(t *testing.T) {
		want, err := FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
		assert.Nil(t, err)
		assert.Equal[string](t, want, "746865206b696420646f6e277420706c6179")
	})
}

func corpusFromFile(name string) map[rune]float64 {
	text, err := os.ReadFile(name)
	if err != nil {
		panic(fmt.Sprintln("failed to read corpus file:", err))
	}

	return buildCorpus(string(text))
}

var corpus = corpusFromFile("../testdata/alice-in-wonderland.txt")

func TestFindSingleByteXORCipher(t *testing.T) {
	res, key, score := FindSingleByteXORCipher([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"), corpus)
	t.Logf("%s %s %v", res, string(key), score)
}
