package signature

import (
	"errors"
	"net/url"
	"fmt"
	"sort"
	"aliyun-acmp-push-go-sdk/models/hmacsha1"
)

func SignatureString(httprequrl *string,httpmethord *string) (signstr *string,err error) {
	if httprequrl==nil {
		return nil,errors.New("SignatureString httprequrl pointer shouldn't be nil")
	}
	u,err:=url.Parse(*httprequrl)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString httprequrl parse error %s",err))
	}
	uParam,err:=url.ParseQuery(u.RawQuery)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString httprequrl ParseQuery error %s",err))
	}
	keys:=make([]string,len(uParam))
	i:=0
	for k,_:=range uParam {
		keys[i]=k
		i++
	}
	sort.Strings(keys)
	querystr:=""
	for _,v:=range keys {
		querystr+="&"+v+"="+uParam[v][0]
	}
	querystr=querystr[1:]
	temp:=url.QueryEscape(querystr)
	return &temp,nil
}

func GetSignature(urlencodestr *string,accesssecret *string) (signstr *string,err error) {
	signstr,err= hmacsha1.GetHmacStr(urlencodestr,accesssecret)
	if err!=nil {
		return nil,errors.New(fmt.Sprint("SignatureString GetHmacStr error %s",err))
	}
	return signstr,nil
}
