package bean

import (
	"encoding/json"
	"errors"
)

type NotifyParam struct {
	Action        *string            `json:"action"`
	AppKey        *string            `json:"app_key"`
	Target        *string            `json:"target"`
	TargetValue   *string            `json:"target_value"`
	Title         *string            `json:"title"`
	Body          *string            `json:"body"`
	ExtParameters *map[string]string `json:"ext_parameters"`
}

func (this *NotifyParam) ToString() (paramstrp *string, err error) {
	if this == nil {
		return nil, errors.New("NotifyParam pointer shouldn't be nil")
	}
	if this.Action == nil || this.AppKey == nil || this.Target == nil || this.TargetValue == nil || this.Title == nil || this.Body == nil {
		return nil, errors.New("NotifyParam some perpoties shouldn't be nil")
	}
	var headstr string
	if this.ExtParameters != nil {
		b, err := json.Marshal(this.ExtParameters)
		if err != nil {
			return nil, err
		}
		headstr += "ExtParameters=" + string(b) + "&"
	}
	headstr += "Action=" + *this.Action + "&AppKey=" + *this.AppKey + "&Target=" + *this.Target + "&TargetValue=" + *this.TargetValue + "&Title=" + *this.Title + "&Body=" + *this.Body
	paramstrp = &headstr
	return paramstrp, nil
}
