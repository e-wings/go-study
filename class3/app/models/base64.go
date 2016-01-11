package models

import (
	"encoding/base64"
)

const encodeTable = "GHIMABUVWCFNOPQRS12TJKLXYZadebcDEfgnopqrstuvhijklmw4567xyz0389-_"

func Base64Encode(str string) string {
	encoding := base64.NewEncoding(encodeTable)
	encodedStr := encoding.EncodeToString([]byte(str))
	return encodedStr
}
func Base64Decode(str string) string {
	encoding := base64.NewEncoding(encodeTable)
	encodedBytes, _ := encoding.DecodeString(str)
	//return fmt.Sprintf("%s", encodedBytes)
	return string(encodedBytes)
}
