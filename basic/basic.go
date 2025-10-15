package basic

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
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
