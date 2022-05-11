package dto

type BaseResponseDto struct {
	Code int32       `json:"code"`
	Data interface{} `json:"data"`
}

func BuildSuccessResp(data interface{}) BaseResponseDto {
	resp := BaseResponseDto{}
	resp.Code = 1
	resp.Data = data
	return resp
}

// BuildResp go 不支持方法重载
func BuildResp() BaseResponseDto {
	return BuildSuccessResp(nil)
}

func BuildFailResp(errMsg string) BaseResponseDto {
	resp := BaseResponseDto{}
	resp.Code = 0
	resp.Data = errMsg
	return resp
}
