package common

type BaseResponseDto struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type GeneralResponse struct {
	Success bool `json:"success"`
}
