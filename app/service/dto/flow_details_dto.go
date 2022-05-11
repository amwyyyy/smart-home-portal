package dto

type FlowDetailsDto struct {
	Value        *FlowDetails `json:"value"`
	ErrorCode    int          `json:"errorCode"`
	ErrorMessage string       `json:"errorMessage"`
}

type FlowDetails struct {
	// 单据id 对应其他api的flowId
	Id string `json:"id"`

	// 是否有效（或者理解为是否被删除） true-有效，false-无效
	Active bool `json:"active"`

	// 创建时间
	CreateTime int64 `json:"createTime"`

	// 更新时间
	UpdateTime int64 `json:"updateTime"`

	// 企业id
	CorporationId string `json:"corporationId"`

	// 流程发起人id
	OwnerId string `json:"ownerId"`

	// 流程发起人默认部门id
	OwnerDefaultDepartment string `json:"ownerDefaultDepartment"`

	// 流程状态 pending-提交中 approving-审批中 rejected-已驳回 paying-待支付 PROCESSING-支付中 paid-已支付 archived-归档 sending-寄送中 receiving-收单中
	State string `json:"state"`

	// 流程类型
	FlowType string `json:"flowType"`

	// 单据类型 expense-报销单  loan-借款单 payment-付款单 requisition-申请单 custom-通用审批单 receipt-收款单
	FormType string `json:"formType"`

	// 单据详情
	Form *Form `json:"form"`
}

type Form struct {
	// 单据编码
	Code string `json:"code"`

	// 单据标题
	Title string `json:"title"`

	// 工号
	WorkCode string `json:"u_工号"`

	// 提交时间
	SubmitDate int64 `json:"submitDate"`

	// 说明
	Description string `json:"description"`

	// 提交人id
	SubmitterId string `json:"submitterId"`

	FeeDatePeriod DatePeriod `json:"feeDatePeriod"`
}

type DatePeriod struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
