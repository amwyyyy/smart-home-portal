package dto

type SysAttendAddReq struct {
	// 考勤来源
	AppName string `json:"appName"`

	// 数据类型
	DataType string `json:"dataType"`

	// 考勤记录 json
	Datas string `json:"datas"`

	// 扩展参数 json
	Others string `json:"others"`
}
