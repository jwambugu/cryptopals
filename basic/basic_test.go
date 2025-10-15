package basic_test

import (
	"cryptopals/assert"
	"cryptopals/basic"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	got, err := basic.HexToBase64(input)
	assert.Nil(t, err)
	assert.Equal[string](t, got, want)
}

func TestFixedXOR(t *testing.T) {
	t.Run("length not equal", func(t *testing.T) {
		want, err := basic.FixedXOR("1c0111001f0101", "1c0111001f01011112")
		assert.Equal[error](t, err, basic.ErrLengthNotEqual)
		assert.True(t, len(want) == 0)
	})

	t.Run("XORs correctly", func(t *testing.T) {
		want, err := basic.FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
		assert.Nil(t, err)
		assert.Equal[string](t, want, "746865206b696420646f6e277420706c6179")
	})
}
