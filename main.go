package main

import (
	"crypto/rand"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/gogf/gf/v2/text/gstr"
	"io"
)

func main() {
	fmt.Println("KEY:", Random(64))
}

func Random(length int) string {
	str := ""

	strLen := 0

	for strLen < length {
		size := length - strLen

		bytes := Bytes(size)

		strBase64 := gbase64.EncodeToString(bytes)

		strReplace := gstr.Replace(strBase64, "/", "")
		strReplace = gstr.Replace(strReplace, "+", "")
		strReplace = gstr.Replace(strReplace, "=", "")

		str += gstr.SubStr(strReplace, 0, size)

		strLen = gstr.LenRune(str)
	}

	return str
}

func Bytes(n int) []byte {
	data := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, data)
	if err != nil {
		panic(err)
	}
	return data
}
