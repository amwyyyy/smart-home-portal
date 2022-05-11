package dto

type KmReviewRequest struct {
	// 文档标题
	DocSubject string
	// 文档模板id
	FdTemplateId string
	// 文档的富文本内容
	DocContent string
	// 流程表单数据
	FormValues FormValue
	// 文档状态，可以为草稿（"10"）或者待审（"20"）两种状态，默认为待审
	DocStatus string
	// 流程发起人，为单值，格式详见人员组织架构的定义说明
	DocCreator string
	// 文档关键字，格式为["关键字1", "关键字2"...]
	FdKeyword string
	// 辅类别，格式为["辅类别1的ID", "辅类别2的ID"...]
	DocProperty string
	// 流程参数
	FlowParam string
}

type FormValue struct {
	FlowId string `json:"fd_flowId,omitempty"`

	NodeId string `json:"fd_nodeId,omitempty"`

	Code string `json:"fd_code,omitempty"`

	Title string `json:"fd_title,omitempty"`

	Id string `json:"fd_id,omitempty"`

	WorkeId string `json:"fd_yuangongid,omitempty"`

	Name string `json:"fd_name,omitempty"`

	RealName string `json:"fd_xingming,omitempty"`

	WorkCode string `json:"fd_gonghao,omitempty"`

	CreateTime string `json:"fd_kaishishijian,omitempty"`

	EndTime string `json:"fd_zhongzhishijian,omitempty"`

	Date string `json:"fd_waichuriqi,omitempty"`

	StartTime string `json:"fd_ks,omitempty"`

	FinishTime string `json:"fd_js,omitempty"`

	OutTime string `json:"fd_ehr_outtime,omitempty"`

	CcTime string `json:"fd_chuchaishichang,omitempty"`

	OtherName string `json:"fd_other_name,omitempty"`

	ErrContent string `json:"fd_content,omitempty"`

	ErrTime string `json:"fd_times,omitempty"`
}
