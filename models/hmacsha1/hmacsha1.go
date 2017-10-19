package hmacsha1

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"encoding/base64"
)

func GetHmacStr(paramstr *string, keystr *string) (hmacstr *string, err error) {
	if paramstr == nil || keystr == nil {
		return nil, errors.New("GetHmacStr parameter pointer shouldn't be nil")
	}
	hmac := hmac.New(sha1.New, []byte(*keystr))
	temp := base64.StdEncoding.EncodeToString(hmac.Sum(nil))
	return &temp, nil
}
