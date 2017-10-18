package acmp

type Notify struct {
	Action        string `json:"action"`
	AppKey        string `json:"app_key"`
	Target        string `json:"target"`
	TargetValue   string `json:"target_value"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	ExtParameters string `json:"ext_parameters"`
}
