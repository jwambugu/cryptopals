package basic

import (
	"bytes"
	"cryptopals/assert"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
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

func readFile(t *testing.T, name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("failed to read file: %s", err)
	}

	return data
}

func TestProblem4(t *testing.T) {
	data := readFile(t, "../testdata/4.txt")

	var (
		bestScore float64
		res       []byte
	)

	for _, line := range strings.Split(string(data), "\n") {
		out, _, score := FindSingleByteXORCipher([]byte(line), corpus)
		if score > bestScore {
			res = out
			bestScore = score
		}
	}

	t.Logf("%s", res)
}

func decodeHex(t *testing.T, s string) []byte {
	t.Helper()

	v, err := hex.DecodeString(s)
	if err != nil {
		t.Fatal("failed to decode hex:", s)
	}

	return v
}

func TestRepeatingXOR(t *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	want := `0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`

	res := RepeatingXOR([]byte(input), []byte("ICE"))
	if !bytes.Equal(res, decodeHex(t, want)) {
		t.Fatal("wrong result:", res)
	}
}
