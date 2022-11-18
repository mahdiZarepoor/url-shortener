package pkg

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5Hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:3])
}