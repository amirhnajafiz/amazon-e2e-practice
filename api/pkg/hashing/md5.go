package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash takes a string input and returns its MD5 hash.
func MD5Hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
