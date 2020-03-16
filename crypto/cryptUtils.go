package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

//GetMd5 from input
func GetMd5(input string) string {
	if input == "" {
		return ""
	}
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}
