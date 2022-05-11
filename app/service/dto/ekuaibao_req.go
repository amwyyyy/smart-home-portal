package dto

type EkuaibaoRequest struct {
	// 单据编码
	Code string `json:"code"`

	// 企业id
	CorporationId string `json:"corporationId"`

	// 单据标题
	Title string `json:"title"`

	// 流程id
	FlowId string `json:"flowId"`

	// 节点id
	NodeId string `json:"nodeId"`

	// 工号
	WorkCode string `json:"u_工号"`

	// 提交者信息
	SubmitterId Submitter `json:"submitterId"`
}

// Submitter 提交者
type Submitter struct {
	Id string `json:"id"`

	Name string `json:"name"`
}