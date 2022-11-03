package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------- Base64Decode ------------------------------------------------------

func TestBase64Decode(t *testing.T) {
	s := "Q0MxMTAwMTEwMA=="
	decodeBytes := Base64Decode([]byte(s))
	exceptBytes := []byte{67, 67, 49, 49, 48, 48, 49, 49, 48, 48, 0, 0}
	assert.Equal(t, exceptBytes, decodeBytes)
}

func ExampleBase64Decode() {
	s := "Q0MxMTAwMTEwMA=="
	decode := Base64Decode([]byte(s))
	fmt.Println(decode)
	// Output:
	// [67 67 49 49 48 48 49 49 48 48 0 0]
}

// ------------------------------------------------- Base64DecodeAsString ----------------------------------------------

func TestBase64DecodeAsString(t *testing.T) {
	s := "Q0MxMTAwMTEwMA=="
	decodeString := Base64DecodeAsString([]byte(s))
	fmt.Println(decodeString)
}

func ExampleBase64DecodeAsString() {
	s := "Q0MxMTAwMTEwMA=="
	decodeString := Base64DecodeAsString([]byte(s))
	fmt.Println(decodeString)
	// Output:
	// CC11001100
}

// ------------------------------------------------- Base64DecodeAsStringE ---------------------------------------------

func TestBase64DecodeAsStringE(t *testing.T) {
	s := "1"
	_, err := Base64DecodeAsStringE([]byte(s))
	assert.NotNil(t, err)
}

func ExampleBase64DecodeAsStringE() {
	s := "Q0MxMTAwMTEwMA=="
	decodeString, err := Base64DecodeAsStringE([]byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(decodeString)
	// Output:
	// CC11001100
}

// ------------------------------------------------- Base64DecodeE -----------------------------------------------------

func TestBase64DecodeE(t *testing.T) {
	s := "Q0MxMTAwMTEwMA=="
	decodeBytes, err := Base64DecodeE([]byte(s))
	assert.Nil(t, err)
	exceptBytes := []byte{67, 67, 49, 49, 48, 48, 49, 49, 48, 48, 0, 0}
	assert.Equal(t, exceptBytes, decodeBytes)
}

func ExampleBase64DecodeE() {
	s := "Q0MxMTAwMTEwMA=="
	decodeBytes, err := Base64DecodeE([]byte(s))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(decodeBytes)
	// Output:
	// [67 67 49 49 48 48 49 49 48 48 0 0]
}

// ------------------------------------------------- Base64Encode ------------------------------------------------------

func TestBase64Encode(t *testing.T) {
	s := "CC11001100"
	encodeBytes := Base64Encode([]byte(s))
	exceptBytes := []byte{81, 48, 77, 120, 77, 84, 65, 119, 77, 84, 69, 119, 77, 65, 61, 61}
	assert.Equal(t, exceptBytes, encodeBytes)
}

func ExampleBase64Encode() {
	s := "CC11001100"
	encodeBytes := Base64Encode([]byte(s))
	fmt.Println(encodeBytes)
	// Output:
	// [81 48 77 120 77 84 65 119 77 84 69 119 77 65 61 61]
}

// ------------------------------------------------- Base64EncodeAsString ----------------------------------------------

func TestBase64EncodeAsString(t *testing.T) {
	s := "CC11001100"
	s2 := Base64EncodeAsString([]byte(s))
	assert.Equal(t, "Q0MxMTAwMTEwMA==", s2)
}

func ExampleBase64EncodeAsString() {
	s := "CC11001100"
	s2 := Base64EncodeAsString([]byte(s))
	fmt.Println(s2)
	// Output:
	// Q0MxMTAwMTEwMA==
}

// ------------------------------------------------- -------------------------------------------------------------------
