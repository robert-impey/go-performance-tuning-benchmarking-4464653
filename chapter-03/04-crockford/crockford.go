package crockford

import (
	"bytes"
	"crypto/rand"
	"encoding/base32"
)

var symbols = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
var encoding = base32.NewEncoding(symbols)
var numBytes = 16
var randBytes = make([]byte, numBytes)

func NewID() string {
	rand.Read(randBytes)
	str := Encode(randBytes)
	return str[:26]
}

func Encode(b []byte) string {
	var buf bytes.Buffer

	enc := base32.NewEncoder(encoding, &buf)
	enc.Write(b)
	enc.Close()
	return buf.String()
}
