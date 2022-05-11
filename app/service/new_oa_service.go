package service

import (
	"bytes"
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/service/dto"
	"com/denis/smarthome/app/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type NewOaService struct {
}

// AddAttend 添加考勤信息
func (e *NewOaService) AddAttend(req dto.SysAttendAddReq) dto.SysAttendResult {
	jsonVal, _ := json.Marshal(req)
	resp, err := http.Post(cfg.NewOaCfg.NewOaBaseUrl+cfg.NewOaCfg.SysAttendUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		logger.Errorf("Error request: %s", err)
	}

	var response dto.SysAttendResult
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.Errorf("Error decode: %s", err)
		return dto.SysAttendResult{}
	}
	logger.Info(utils.ToString(response))

	return response
}

// AddReview 添加流程审批
func (e *NewOaService) AddReview(req dto.KmReviewRequest) error {
	values := url.Values{}
	values.Set("docSubject", req.DocSubject)
	values.Set("fdTemplateId", req.FdTemplateId)
	values.Set("docContent", req.DocContent)
	values.Set("formValues", utils.ToString(req.FormValues))
	values.Set("docStatus", "20")
	values.Set("docCreator", "{\"PersonNo\": \""+req.DocCreator+"\"}")

	logger.Info("addReview request: ", utils.ToString(values))
	resp, err := http.PostForm(cfg.NewOaCfg.NewOaBaseUrl+cfg.NewOaCfg.KmReviewUrl, values)
	if err != nil {
		logger.Errorf("Error request: %s", err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("Error read: %s", err)
		return err
	}

	bodyStr := utils.ToString(body)
	logger.Info("addReview: ", bodyStr)

	if len(bodyStr) != 44 {
		return errors.New(bodyStr + "@@@" + utils.ToString(values))
	}

	return nil
}

// CalOutTime 计算外出时间
func (e *NewOaService) CalOutTime(empCode string, beginDate string, endDate string, attendType string) (float64, error) {
	params := map[string]string{}
	params["vKey"] = "b23bfe74d884cedec818e2e57349c64f"
	params["attendanceItem"] = attendType
	params["empCode"] = empCode
	params["beginDate"] = beginDate
	params["endDate"] = endDate

	jsonVal, _ := json.Marshal(params)
	resp, err := http.Post(cfg.NewOaCfg.EhrUrl, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		logger.Errorf("Error request: %s", err)
		return 0.0, err
	}

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.Errorf("Error decode: %s", err)
		return 0.0, err
	}

	if response["errcode"] != "0" {
		logger.Errorf("calOutTime err, empCode: %s, beginDate: %s, endDate: %s, err: %s", empCode, beginDate, endDate, response["errmsg"])
		return 0.0, nil
	}

	return response["result"].(float64), nil
}

func (e *NewOaService) CheckTime(flowId string) error {
	transformService := TransformService{}
	flowDetail := transformService.GetFlowDetails(flowId)
	if flowDetail.ErrorCode == 0 {
		if flowDetail.Value.Form.FeeDatePeriod.Start > flowDetail.Value.Form.FeeDatePeriod.End {
			logger.Infof("time format error, start: %d, end: %d", flowDetail.Value.Form.FeeDatePeriod.Start, flowDetail.Value.Form.FeeDatePeriod.End)
			return errors.New("请选择正确的开始时间与结束时间")
		}

		//// 出差
		//if helper.GetFlowType(flowDetail.Value.Form.Title) == helper.BusinessType {
		//	start := utils.FormatForLayout(flowDetail.Value.Form.FeeDatePeriod.Start, utils.Short)
		//	end := utils.FormatForLayout(flowDetail.Value.Form.FeeDatePeriod.End, utils.Short)
		//	result, err := e.CalOutTime(flowDetail.Value.Form.WorkCode, start, end, cfg.NewOaCfg.BtAttendType)
		//	if err == nil {
		//		if result <= 0 {
		//			logger.Infof("time format error, start: %s, end: %s", start, end)
		//			return errors.New("请选择正确的开始时间与结束时间")
		//		}
		//	}
		//} else {
		//	// 外出
		//	// 只检查单天的
		//	if utils.IsSameDay(flowDetail.Value.Form.FeeDatePeriod.Start, flowDetail.Value.Form.FeeDatePeriod.End) {
		//		start := utils.FormatForLayout(flowDetail.Value.Form.FeeDatePeriod.Start, utils.Normal)
		//		end := utils.FormatForLayout(flowDetail.Value.Form.FeeDatePeriod.End, utils.Normal)
		//		result, err := e.CalOutTime(flowDetail.Value.Form.WorkCode, start, end, cfg.NewOaCfg.GoAttendType)
		//		if err == nil {
		//			if result < 1 {
		//				logger.Infof("time format error, start: %s, end: %s", start, end)
		//				return errors.New("请选择正确的开始时间与结束时间")
		//			}
		//		}
		//	}
		//}
	}

	return nil
}
