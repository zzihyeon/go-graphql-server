package types

type StandardResponse struct {
	Code    int             `json:"code"`
	Payload StandardPayload `json:"payload"`
}

type StandardPayload struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
