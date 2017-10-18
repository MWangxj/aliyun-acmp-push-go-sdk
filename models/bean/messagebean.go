package bean

import "errors"

type MessageParam struct {
	Action      *string `json:"action"`
	AppKey      *string `json:"app_key"`
	Target      *string `json:"target"`
	TargetValue *string `json:"target_value"`
	Title       *string `json:"title"`
	Body        *string `json:"body"`
}

func (this *MessageParam) ToString() (paramstrp *string, err error) {
	if this == nil {
		return nil, errors.New("MessageParam pointer shouldn't be nil")
	}
	if this.Action == nil || this.AppKey == nil || this.Target == nil || this.TargetValue == nil || this.Title == nil || this.Body == nil {
		return nil, errors.New("MessageParam some perpoties shouldn't be nil")
	}
	var headstr string
	headstr += "Action=" + *this.Action + "&AppKey=" + *this.AppKey + "&Target=" + *this.Target + "&TargetValue=" + *this.TargetValue + "&Title=" + *this.Title + "&Body=" + *this.Body
	paramstrp = &headstr
	return paramstrp, nil
}
