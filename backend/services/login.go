package services

type LoginResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Name   string `json:"name"`
		Cookie string `json:"cookie"`
	}
}
