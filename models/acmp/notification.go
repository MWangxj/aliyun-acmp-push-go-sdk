package acmp

import (
	"giter.org/duoxieyun/openapi/controllers/com"
)


const url  ="http://cloudpush.aliyuncs.com/"

func SendNotifyByAcmp()  {
	
}

func sendNotifyToAndroid(notify *Notify){
	if notify==nil ||notify.Action==""{
		return
	}
	strToSign:="GET/&Action="+notify.Action+"&AppKey="+notify.AppKey+"&Target="+notify.Target+
		"&TargetValue="+notify.TargetValue+"&Title="+notify.Title+"&Body="+notify.Body+"&ExtParameters"+notify.ExtParameters
	urlencodestr:=com.UrlEncode(strToSign)
	println(urlencodestr)
	//httpurl:=url+"?Action="+notify.Action+"&AppKey="+notify.AppKey+"&Target="+notify.Target+
	//	"&TargetValue="+notify.TargetValue+"&Title="+notify.Title+"&Body="+notify.Body+"&ExtParameters"+notify.ExtParameters
	//response,err:=http.Get(httpurl)
	//if err!=nil {
	//	log.GLog.Error("[sendNotifyToAndroid] http.Get err: %s", err)
	//}
}

//http://cloudpush.aliyuncs.com/?Action=PushNoticeToAndroid&AppKey=23267207&Target=ALL&TargetValue=ALL&Title=title&Body=body
