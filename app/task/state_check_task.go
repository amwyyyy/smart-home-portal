package task

import (
	"com/denis/smarthome/app/config"
	"com/denis/smarthome/app/service"
)

var cfg = config.Cfg
var newOaService = service.NewOaService{}

// Task 定时根据flowId获取单据状态
//func Task() {
//	c := cron.New(cron.WithSeconds())
//
//	transformService := service.TransformService{}
//
//	_, err := c.AddFunc("0 0/5 * * * ?", func() {
//		logger.Info("task run!")
//
//		for _, item := range utils.List() {
//			flowDetail := transformService.GetFlowDetails(item)
//
//			if flowDetail.ErrorCode == 0 {
//				logger.Info("flowId:", item, "state:", flowDetail.Value.State)
//				switch flowDetail.Value.State {
//				case "pending":
//				case "approving":
//				case "rejected":
//				case "paying":
//				case "PROCESSING":
//				case "paid":
//					err := AddReview(flowDetail)
//					if err == nil {
//						utils.Remove(item)
//					} else {
//						// 异常单处理
//						errReview(flowDetail.Value.Form.Title, err.Error())
//
//						// 对特殊处理的异常单移除，不要影响原来逻辑
//						if strings.Contains(err.Error(), "@@@") {
//							utils.Remove(item)
//						}
//					}
//				case "archived":
//					logger.Info("state archived, remove flowId:", item)
//					utils.Remove(item)
//				case "sending":
//				case "receiving":
//				default:
//					logger.Info("no support state:", flowDetail.Value.State, "flowId:", item)
//					utils.Remove(item)
//				}
//			} else if flowDetail.ErrorCode == 412 {
//				// 412 单据不存在，则删除item
//				logger.Info("remove 412 flowId:", item)
//				utils.Remove(item)
//			} else {
//				logger.Info("get flow details fail, errorCode:", flowDetail.ErrorCode, ", errMsg:", flowDetail.ErrorMessage)
//			}
//		}
//	})
//	if err != nil {
//		logger.Errorf("task run error:", err)
//	}
//
//	c.Start()
//}
//
//// AddReview oa增加考勤记录
//func AddReview(flowDetails dto.FlowDetailsDto) error {
//	flowType := helper.GetFlowType(flowDetails.Value.Form.Title)
//	if flowType == helper.BusinessType {
//		// 增加出差申请
//		logger.Info("addReview flowType: ", flowType, ", flowId: ", flowDetails.Value.Id)
//		return oaReview(flowDetails, func(formValue *dto.FormValue) error {
//			formValue.CreateTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.Start, utils.Short)
//			formValue.EndTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.End, utils.Short)
//			formValue.CcTime = getWorkTime(flowDetails.Value.Form.FeeDatePeriod.Start, flowDetails.Value.Form.FeeDatePeriod.End)
//			return nil
//		})
//	} else {
//		// 增加外出申请
//		return addGoReview(flowDetails)
//	}
//}
//
//// 增加外出申请
//func addGoReview(flowDetails dto.FlowDetailsDto) error {
//	// 判断是否只外出一天
//	if utils.IsSameDay(flowDetails.Value.Form.FeeDatePeriod.Start, flowDetails.Value.Form.FeeDatePeriod.End) {
//		logger.Info("addReview flowType: ", helper.GoType, ", flowId: ", flowDetails.Value.Id, ", one day")
//		return oaReview(flowDetails, func(formValue *dto.FormValue) error {
//			formValue.CreateTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.Start, utils.Time)
//			formValue.EndTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.End, utils.Time)
//			formValue.Date = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.Start, utils.Short)
//			formValue.StartTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.Start, utils.Normal)
//			formValue.FinishTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.End, utils.Normal)
//			formValue.OutTime = getWorkTime(flowDetails.Value.Form.FeeDatePeriod.Start, flowDetails.Value.Form.FeeDatePeriod.End)
//			return nil
//		})
//	}
//
//	// 外出多天，需要拆成每天提交一张单
//	dates := utils.CalDateGap(flowDetails.Value.Form.FeeDatePeriod.Start, flowDetails.Value.Form.FeeDatePeriod.End)
//	logger.Info("addReview flowType: ", helper.GoType, ", flowId: ", flowDetails.Value.Id, ", ", len(dates), " day")
//	for index, date := range dates {
//		err := oaReview(flowDetails, func(formValue *dto.FormValue) error {
//			formValue.Date = date
//
//			// 对拆分单的时间处理，首尾取实际时间，中间使用早9晚18
//			if index == 0 {
//				formValue.CreateTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.Start, utils.Time)
//				formValue.EndTime = "18:00"
//			} else if index == len(dates)-1 {
//				formValue.CreateTime = "09:00"
//				formValue.EndTime = utils.FormatForLayout(flowDetails.Value.Form.FeeDatePeriod.End, utils.Time)
//			} else {
//				formValue.CreateTime = "09:00"
//				formValue.EndTime = "18:00"
//			}
//
//			formValue.StartTime = formValue.Date + " " + formValue.CreateTime
//			formValue.FinishTime = formValue.Date + " " + formValue.EndTime
//			formValue.OutTime = getWorkTime(flowDetails.Value.Form.FeeDatePeriod.Start, flowDetails.Value.Form.FeeDatePeriod.End)
//
//			return nil
//		})
//		if err != nil {
//			logger.Errorf("addReview error: ", err.Error(), ", date: ", date)
//			return err
//		}
//	}
//
//	return nil
//}
//
//func getWorkTime(startTime int64, finishTime int64) string {
//	return strconv.Itoa(int(utils.CalTimeDiff(startTime, finishTime)))
//}
//
//// GetTemplateId 获取模板id
//func getTemplateId(title string) string {
//	if helper.GetFlowType(title) == helper.BusinessType {
//		return cfg.NewOaCfg.BtFdTemplateId
//	} else {
//		return cfg.NewOaCfg.GoFdTemplateId
//	}
//}
//
//func oaReview(flowDetails dto.FlowDetailsDto, apply func(formValue *dto.FormValue) error) error {
//	req := dto.KmReviewRequest{}
//	req.FdTemplateId = getTemplateId(flowDetails.Value.Form.Title)
//	req.DocSubject = flowDetails.Value.Form.Title
//	req.DocCreator = helper.WorkCodeTransform(flowDetails.Value.Form.WorkCode)
//
//	formValue := dto.FormValue{}
//	formValue.FlowId = flowDetails.Value.Id
//	formValue.Id = flowDetails.Value.Id
//	formValue.Code = flowDetails.Value.Form.Code
//	formValue.Title = flowDetails.Value.Form.Title
//	formValue.Name = helper.GetName(flowDetails.Value.Form.Title)
//	formValue.RealName = formValue.Name
//	formValue.WorkeId = flowDetails.Value.Form.SubmitterId
//	formValue.WorkCode = flowDetails.Value.Form.WorkCode
//	err := apply(&formValue)
//	if err != nil {
//		return err
//	}
//
//	req.FormValues = formValue
//	err = newOaService.AddReview(req)
//	return err
//}
//
//func errReview(title string, errMsg string) {
//	req := dto.KmReviewRequest{}
//	req.FdTemplateId = cfg.NewOaCfg.ErrTemplateId
//	req.DocSubject = "接口异常告警单"
//	req.DocCreator = "API0001"
//
//	name := "出差接口数据对接单（勿人工提单）"
//	if helper.GetFlowType(title) == helper.GoType {
//		name = "外出接口数据对接单（勿人工提单）"
//	}
//
//	content := errMsg
//	if strings.Contains(errMsg, "@@@") {
//		msgs := strings.Split(errMsg, "@@@")
//
//		sDec, err := base64.StdEncoding.DecodeString(msgs[0])
//		if err != nil {
//			name = name + " - " + msgs[0]
//		} else {
//			name = name + " - " + string(sDec)
//		}
//
//		content = msgs[1]
//	}
//
//	formValue := dto.FormValue{}
//	formValue.Name = name
//	formValue.ErrContent = content
//	formValue.ErrTime = time.Now().Format("2006-01-02 15:04:05")
//
//	req.FormValues = formValue
//	err := newOaService.AddReview(req)
//	if err != nil {
//		logger.Errorf("Error: %s", err)
//	}
//}
