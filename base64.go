package slice

import "encoding/base64"

// ------------------------------------------------ ---------------------------------------------------------------------

// Base64Encode 把字节数组编码为base64的字节数组
func Base64Encode(byteSlice []byte) []byte {
	if len(byteSlice) == 0 {
		return nil
	}
	dest := make([]byte, base64.StdEncoding.EncodedLen(len(byteSlice)))
	base64.StdEncoding.Encode(dest, byteSlice)
	return dest
}

// Base64EncodeAsString 把字节数组编码为base64字符串
func Base64EncodeAsString(byteSlice []byte) string {
	if len(byteSlice) == 0 {
		return ""
	}
	return string(Base64Encode(byteSlice))
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Base64Decode 把Base64编码的字节数组解析为字节数组，如果发生了错误则忽略错误并返回一个空字节数组
func Base64Decode(byteSlice []byte) []byte {
	if len(byteSlice) == 0 {
		return nil
	}
	result, err := Base64DecodeE(byteSlice)
	if err != nil {
		return []byte{}
	}
	return result
}

// Base64DecodeE 把Base64编码的字节数组解码，如果发生了错误则返回错误
func Base64DecodeE(byteSlice []byte) ([]byte, error) {
	if len(byteSlice) == 0 {
		return nil, nil
	}
	dest := make([]byte, base64.StdEncoding.DecodedLen(len(byteSlice)))
	_, err := base64.StdEncoding.Decode(dest, byteSlice)
	if err != nil {
		return nil, err
	}
	return dest, nil
}

// Base64DecodeAsStringE 把Base64编码的字节数组解码为字符串，如果发生了错误则返回错误
func Base64DecodeAsStringE(byteSlice []byte) (string, error) {
	if len(byteSlice) == 0 {
		return "", nil
	}
	bytes, err := Base64DecodeE(byteSlice)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Base64DecodeAsString 把Base64编码的字节数组解码为字符串，如果发生了错误则忽略错误返回空字符串
func Base64DecodeAsString(byteSlice []byte) string {
	if len(byteSlice) == 0 {
		return ""
	}
	result, err := Base64DecodeAsStringE(byteSlice)
	if err != nil {
		return ""
	}
	return result
}

// ------------------------------------------------- ------------------------------------------------------------------------

// TODO
//// BinaryToAscii 把二进制转为ascii字符串，用于兼容浏览器环境中的btoa
//func BinaryToAscii(binaryBase64String string) string {
//
//}
//
// TODO
//// AsciiToBinary 把ascii转为二进制，用于兼容浏览器环境中的atob
//func AsciiToBinary(asciiString string) []byte {
//
//}

// ------------------------------------------------- ------------------------------------------------------------------------
