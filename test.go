package main

import (
	"fmt"
	"aliyun-acmp-push-go-sdk/models/signature"
)

func main() {
	/*
	str:="abcde"
	var strp *string
	strp=&str
	keys:="cdefg"
	var keysp *string
	keysp=&keys
	strr,err:=hmacsha1.GetHmacStr(strp,keysp)
	if err==nil {
		fmt.Print(*strr)
	}*/

	httpurl:="http://ecs.aliyuncs.com/?TimeStamp=2016-02-23T12:46:24Z&Format=XML&AccessKeyId=testid&Action=DescribeRegions&SignatureMethod=HMAC-SHA1&SignatureNonce=3ee8c1b8-83d3-44af-a94f-4e0ad82fd6cf&Version=2014-05-26&SignatureVersion=1.0"
	var httpurlp *string
	httpurlp=&httpurl
	httpmethod:="GET"
	var httpmethodp *string
	httpmethodp=&httpmethod
	signstr,err:=signature.SignatureString(httpurlp,httpmethodp)
	if err==nil {
		fmt.Println(*signstr)
	}

	str:="GET&%2F&AccessKeyId%3Dtestid&Action%3DDescribeRegions&Format%3DXML&SignatureMethod%3DHMAC-SHA1&SignatureNonce%3D3ee8c1b8-83d3-44af-a94f-4e0ad82fd6cf&SignatureVersion%3D1.0&TimeStamp%3D2016-02-23T12%253A46%253A24Z&Version%3D2014-05-26"
	fmt.Println(str)
	var strp *string
	strp=&str
	keys:="testsecret&"
	var keysp *string
	keysp=&keys
	signstrr,err:=signature.GetSignature(signstr,keysp)
	strpr,_:=signature.GetSignature(strp,keysp)
	if err==nil {
		fmt.Println(*signstrr)
		fmt.Println(*strpr)
	}


}
