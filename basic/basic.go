package basic

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"unicode/utf8"
)

var ErrLengthNotEqual = fmt.Errorf("length not equal")

// HexToBase64 converts a hexadecimal-encoded string to its Base64 representation.
func HexToBase64(input string) (string, error) {
	hexBytes, err := hex.DecodeString(input)
	if err != nil {
		return "", fmt.Errorf("decode hex string: %w", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(hexBytes)
	return base64Str, nil
}

// FixedXOR takes two equal-length hexadecimal-encoded strings, decodes them,
// and returns their bitwise XOR as a hexadecimal-encoded string.
//
// If the input strings are of unequal length, it returns ErrLengthNotEqual.
func FixedXOR(inputOne, inputTwo string) (string, error) {
	if len(inputOne) != len(inputTwo) {
		return "", ErrLengthNotEqual
	}

	inputOneBytes, err := hex.DecodeString(inputOne)
	if err != nil {
		return "", fmt.Errorf("decode input one hex string: %w", err)
	}

	inputTwoBytes, err := hex.DecodeString(inputTwo)
	if err != nil {
		return "", fmt.Errorf("decode input two hex string: %w", err)
	}

	out := make([]byte, len(inputOneBytes))

	for i := 0; i < len(inputOneBytes); i++ {
		out[i] = inputOneBytes[i] ^ inputTwoBytes[i]
	}

	return hex.EncodeToString(out), nil
}

// buildCorpus constructs a normalized frequency distribution (corpus)
// from the provided reference text.
//
// It counts how often each rune appears and then divides by the total number
// of runes to get the relative frequency of each character.
func buildCorpus(text string) map[rune]float64 {
	c := make(map[rune]float64)

	for _, char := range text {
		c[char]++
	}

	total := utf8.RuneCountInString(text)

	for char := range c {
		c[char] = c[char] / float64(total)
	}

	return c
}

// scoreEnglish evaluates how closely a piece of text matches the given frequency corpus.
//
// It computes a heuristic "English-likeness" score by summing the frequency
// weights of each rune (from the provided corpus) found in the text, then normalizing
// by the text length. The higher the returned value, the more the text resembles English.
func scoreEnglish(text string, corpus map[rune]float64) float64 {
	var score float64

	for _, char := range text {
		score += corpus[char]
	}

	return score / float64(utf8.RuneCountInString(text))
}

// singleXOR applies a single-byte XOR cipher to the input bytes using the given key.
//
// Each byte of the input slice is XORed with the key to produce a new slice.
func singleXOR(in []byte, key byte) []byte {
	res := make([]byte, len(in))
	for i, c := range in {
		res[i] = c ^ key
	}

	return res
}

// FindSingleByteXORCipher attempts to break a single-byte XOR cipher by brute-forcing all 256 keys.
//
// For each possible key, it decrypts the input using singleXOR(), then scores the result
// using scoreEnglish() against the given corpus. The best-scoring plaintext and its key
// are returned.
func FindSingleByteXORCipher(in []byte, c map[rune]float64) (res []byte, key byte, score float64) {
	for k := 0; k < 256; k++ {
		out := singleXOR(in, byte(k))
		s := scoreEnglish(string(out), c)
		if s > score {
			res = out
			score = s
			key = byte(k)
		}
	}

	return
}
