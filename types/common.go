package types

type JHStockMSStandardResponse struct {
	Code    int             `json:"code"`
	Payload StandardPayload `json:"payload"`
}

type StandardPayload struct {
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
