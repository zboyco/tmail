package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand/v2"
	"unsafe"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func RandomStr(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.IntN(len(charset))]
	}
	return Bytes2Str(b)
}

func Str2Bytes(str string) []byte {
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func Bytes2Str(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
