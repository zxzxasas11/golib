package resp

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Page struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}
