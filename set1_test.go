package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ConvertHexToBase64(t *testing.T) {
	var (
		input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		want  = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	)

	b, err := hex.DecodeString(input)
	require.NoError(t, err)

	got := base64.StdEncoding.EncodeToString(b)
	require.Equal(t, want, got)
}

func Test_Fixed_XOR(t *testing.T) {
	var (
		inputOne = "1c0111001f010100061a024b53535009181c"
		inputTwo = "686974207468652062756c6c277320657965"
		want     = "746865206b696420646f6e277420706c6179"
	)

	inputOneBytes, err := hex.DecodeString(inputOne)
	require.NoError(t, err)

	inputTwoBytes, err := hex.DecodeString(inputTwo)
	require.NoError(t, err)

	out := make([]byte, len(inputOneBytes))
	for i := 0; i < len(inputOneBytes); i++ {
		out[i] = inputOneBytes[i] ^ inputTwoBytes[i]
	}

	got := hex.EncodeToString(out)
	require.Equal(t, want, got)
}
