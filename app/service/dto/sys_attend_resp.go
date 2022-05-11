package dto

type SysAttendResult struct {
	// 0:表示未操作
	// 1:表示操作失败
	// 2:表示操作成功
	ReturnState int32 `json:"returnState"`

	// 返回状态值为0时，该值返回空。
	// 返回状态值为1时，该值错误信息。
	// 返回状态值为2时， 该值返回空。
	Message string `json:"message"`
}
