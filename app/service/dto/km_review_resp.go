package dto

type KmReviewResponse struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Msg     string `json:"msg"`
	Code    string `json:"code"`
}
