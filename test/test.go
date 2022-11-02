package main

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"strings"
)

func StrPad(input string, padLength int, padStr ...string) string {
	output := ""
	inputLen := len(input)
	if inputLen >= padLength {
		return input
	}
	padString := " "
	if len(padStr) > 0 {
		padString = padStr[0]
	}
	padStringLen := len(padString)
	needFillLen := padLength - inputLen
	if diffLen := padStringLen - needFillLen; diffLen > 0 {
		padString = padString[diffLen:]
	}
	for i := 1; i <= needFillLen; i += padStringLen {
		output += padString
	}
	return input + output
}

func intToBytes(n int64, k int) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)

	gbyte := bytesBuffer.Bytes()
	x := len(gbyte)
	nb := make([]byte, k)
	for i := 0; i < k; i++ {
		nb[i] = gbyte[x-i-1]
	}
	return nb
}

func main() {
	type data struct {
		uid   string
		sid   string
		token string
		uname string
		sip   string
		ctime int64
		ltime int64
	}
	d := data{
		uid:   "1",
		sid:   "fdfdsafda",
		token: "fafdafdsaf",
		uname: "test@imooc.com",
		sip:   "10.98.16.63",
		ctime: 1002003001,
		ltime: 1002003001,
	}
	str := fmt.Sprintf("%s%s%s%s%s%s%s",
		StrPad(d.sid, 64),
		StrPad(d.sip, 32),
		StrPad(d.uid, 32),
		StrPad(d.uname, 32),
		StrPad(d.token, 32),
		intToBytes(d.ctime, 4),
		intToBytes(d.ltime, 4),
	)
	fmt.Println(str)
	apsid := base64.StdEncoding.EncodeToString([]byte(str))
	apsid = strings.ReplaceAll(apsid, "I", "A")
	apsid = strings.ReplaceAll(apsid, "C", "A")
	apsid = strings.ReplaceAll(apsid, "g", "A")
	fmt.Println(apsid)
}
