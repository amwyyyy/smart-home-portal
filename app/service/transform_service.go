package service

import (
	"bytes"
	"com/denis/smarthome/app/config"
	"com/denis/smarthome/app/helper"
	"com/denis/smarthome/app/logger"
	"com/denis/smarthome/app/service/dto"
	"com/denis/smarthome/app/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type TransformService struct {
}

var accessToken = dto.GetAccessTokenResp{}
var cfg = config.Cfg

// GetAccessToken 获取授权码
func (e *TransformService) GetAccessToken(force bool) string {
	if force || (time.Now().Unix()+10000) > accessToken.Value.ExpireTime {
		body := map[string]string{"appKey": cfg.EkuaibaoCfg.AppKey, "appSecurity": cfg.EkuaibaoCfg.AppSecurity}
		jsonVal, _ := json.Marshal(body)

		resp, err := http.Post(cfg.EkuaibaoCfg.ServicePath+cfg.EkuaibaoCfg.Api.GetAccessTokenURI, "application/json", bytes.NewBuffer(jsonVal))
		if err != nil {
			logger.Errorf("Error request: %s", err)
			return ""
		}
		defer func() {
			_ = resp.Body.Close()
		}()

		var response dto.GetAccessTokenResp
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			helper.Panicf("Error decode: %s", err)
		}

		logger.Info("GetAccessToken: ", utils.ToString(response))

		accessToken.Value.AccessToken = response.Value.AccessToken
		accessToken.Value.ExpireTime = response.Value.ExpireTime
	}

	return accessToken.Value.AccessToken
}

// GetFlowDetails 获取单据详情
func (e *TransformService) GetFlowDetails(flowId string) dto.FlowDetailsDto {
	var response dto.FlowDetailsDto

	accessToken := e.GetAccessToken(false)
	if accessToken == "" {
		response.ErrorCode = 555
		response.ErrorMessage = "get access token error!"
		return response
	}

	reqUrl, err := url.Parse(cfg.EkuaibaoCfg.ServicePath + cfg.EkuaibaoCfg.Api.FlowDetailsURI)
	if err != nil {
		logger.Errorf("Error servicePath: %s", err)
		response.ErrorCode = 500
		response.ErrorMessage = "config error!"
		return response
	}

	params := url.Values{}
	params.Set("accessToken", accessToken)
	params.Set("flowId", flowId)

	reqUrl.RawQuery = params.Encode()
	urlPath := reqUrl.String()

	resp, err := http.Get(urlPath)
	if err != nil {
		logger.Errorf("Error requestUrl: %s, error: %s", urlPath, err)
		response.ErrorCode = 556
		response.ErrorMessage = "network error!"
		return response
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		response.ErrorCode = resp.StatusCode
		if b, err := ioutil.ReadAll(resp.Body); err == nil {
			response.ErrorMessage = string(b)
		}
		logger.Errorf("Error flowId: %s, resp: %s", flowId, utils.ToString(response))
		e.GetAccessToken(true)
		return response
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		logger.Errorf("Error flowId: %s, resp: %s", flowId, err)
	}
	return response
}
