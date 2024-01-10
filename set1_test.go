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
