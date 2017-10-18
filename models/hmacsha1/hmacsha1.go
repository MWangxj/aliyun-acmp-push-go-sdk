package hmacsha1

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

func GetHmacStr(paramstr *string, keystr *string) (hmacstr *string, err error) {
	if paramstr == nil || keystr == nil {
		return nil, errors.New("GetHmacStr parameter pointer shouldn't be nil")
	}
	hmac := hmac.New(sha1.New, []byte(*keystr))
	temp := hex.EncodeToString(hmac.Sum(nil))
	return &temp, nil
}
