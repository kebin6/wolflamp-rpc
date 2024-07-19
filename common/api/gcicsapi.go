package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/kebin6/wolflamp-rpc/common/enum/cachekey"
	"github.com/kebin6/wolflamp-rpc/common/util"
	"github.com/kebin6/wolflamp-rpc/ent/exchange"
	"github.com/kebin6/wolflamp-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	StatusSuccess = 1
	StatusFail    = 0
	StatusUnAuth  = 1001
)

type BaseResponse struct {
	Status int64  `json:"status"`
	Code   int64  `json:"code"`
	Msg    string `json:"msg"`
}

type ErrorInfo struct {
	ErrorCode int64 `json:"errCode"`
}

type BalanceResponse struct {
	BaseResponse
	Data BalanceData `json:"data"`
}

type BalanceData struct {
	ErrorInfo
	Token *float64 `json:"token"`
	Coin  *float64 `json:"coin"`
}

type GcicsApi struct {
	Ctx       context.Context
	SvcCtx    *svc.ServiceContext
	UserId    uint64
	GameToken string
}

func (g GcicsApi) CheckStatus(status int64, code int64, msg string) error {
	switch status {
	case StatusUnAuth:
		// 先删除旧token
		err := g.SvcCtx.Redis.Del(g.Ctx, fmt.Sprintf(string(cachekey.GameAuthToken), g.UserId)).Err()
		if err != nil {
			logx.Errorw("failed to delete token in redis", logx.Field("detail", err))
			return errorx.NewCodeInternalError(fmt.Sprintf("failed[%d:%d]:%s", status, code, msg))
		}
		return errorx.NewCodeError(http.StatusUnauthorized, "Unauthorized")
	case StatusSuccess:
		return nil
	default:
		return errorx.NewCodeInternalError(fmt.Sprintf("failed[%d:%d]:%s", status, code, msg))
	}
}

func (g GcicsApi) GetBalance(gameToken string) (*BalanceData, error) {
	c := g.SvcCtx.Config.GcicsConf
	postUrl := c.Host + "/Game/api/wallet"
	timestamp := time.Now().Unix()
	sign, err := util.GenerateSHA1Signature(strconv.FormatInt(timestamp, 10), c.AppId, c.AppSecret)
	if err != nil {
		return nil, err
	}

	// 创建表单数据
	formData := url.Values{}
	// 将表单数据转换为字节流
	bodyReader := bytes.NewBufferString(formData.Encode())
	// 创建请求
	req, err := http.NewRequest("POST", postUrl, bodyReader)
	if err != nil {
		return nil, err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-PD-SIGN", sign)
	req.Header.Set("GameToken", gameToken)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Failed to get balance[%d]", resp.StatusCode))
	}

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Error reading response body: %s", err.Error()))
	}

	// 解析JSON响应到BaseResponse结构体
	var postResp BalanceResponse
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}

	if err = g.CheckStatus(postResp.Status, postResp.Code, postResp.Msg); err != nil {
		return nil, err
	}
	return &postResp.Data, nil
}

type GeneratePaymentLinkReq struct {
	OrderId    string  `json:"orderid"`
	Ref        string  `json:"ref"`
	Coin       string  `json:"coin"`
	Amount     float64 `json:"amount"`
	Time       int64   `json:"time"`
	NotifyUrl  string  `json:"notify_url"`
	SuccessUrl string  `json:"success_url"`
	FailUrl    string  `json:"fail_url"`
}

type GeneratePaymentLinkResp struct {
	BaseResponse
	Data string `json:"data"`
	Link string `json:"url"`
}

// GeneratePaymentLink 请求生成支付链接
func (g GcicsApi) GeneratePaymentLink(exchangeId uint64, coinType string, amount float64) (*GeneratePaymentLinkResp, error) {
	c := g.SvcCtx.Config.GcicsConf
	postUrl := c.Host + "/Game/api/generateCustomLink"
	timestamp := time.Now().Unix()
	sign, err := util.GenerateSHA1Signature(strconv.FormatInt(timestamp, 10), c.AppId, c.AppSecret)
	if err != nil {
		return nil, err
	}

	// 定义要发送的数据
	data := &GeneratePaymentLinkReq{
		OrderId:    strconv.FormatUint(exchangeId, 10),
		Ref:        "Game",
		Coin:       coinType,
		Amount:     amount,
		Time:       timestamp,
		NotifyUrl:  c.NotifyUrl,
		SuccessUrl: fmt.Sprintf(c.SuccessUrl, exchangeId),
		FailUrl:    fmt.Sprintf(c.FailUrl, exchangeId),
	}

	fmt.Printf("send GeneratePaymentLinkReq: notifyurl=%s", c.NotifyUrl)
	// 将数据转换为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 创建请求
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-PD-SIGN", sign)
	req.Header.Set("APP-ID", c.AppId)
	req.Header.Set("GameToken", g.GameToken)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Failed to get payment link[%d]", resp.StatusCode))
	}

	// 解析JSON响应到BaseResponse结构体
	var basePesp BaseResponse
	err = json.Unmarshal(body, &basePesp)
	if err != nil {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}
	if err = g.CheckStatus(basePesp.Status, basePesp.Code, basePesp.Msg); err != nil {
		return nil, err
	}
	var postResp GeneratePaymentLinkResp
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return nil, errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}

	err = g.SvcCtx.DB.Exchange.Update().
		Where(exchange.ID(exchangeId)).
		SetGcicsOrderID(postResp.Data).
		Exec(g.Ctx)
	if err != nil {
		return nil, err
	}
	return &postResp, nil
}

type WithdrawReq struct {
	Coin   string  `json:"coin"`
	Amount float64 `json:"amount"`
	Time   int64   `json:"time"`
}

type WithdrawResp struct {
	BaseResponse
}

func (g GcicsApi) Withdraw(coinType string, lambAmount float64) error {
	c := g.SvcCtx.Config.GcicsConf
	postUrl := c.Host + "/Game/api/withdrawal"
	timestamp := time.Now().Unix()
	sign, err := util.GenerateSHA1Signature(strconv.FormatInt(timestamp, 10), c.AppId, c.AppSecret)
	if err != nil {
		return err
	}

	// 定义要发送的数据
	data := &WithdrawReq{
		Coin:   coinType,
		Amount: lambAmount,
		Time:   timestamp,
	}

	// 将数据转换为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 创建请求
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-PD-SIGN", sign)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errorx.NewCodeInternalError(fmt.Sprintf("Failed to withdraw[%d]", resp.StatusCode))
	}

	// 解析JSON响应到BaseResponse结构体
	var postResp GeneratePaymentLinkResp
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}

	return g.CheckStatus(postResp.Status, postResp.Code, postResp.Msg)
}

type CommissionInfo struct {
	UserId     uint64  `json:"user_id"`
	LambAmount float64 `json:"amount"`
}
type CommissionReq struct {
	Coin string           `json:"coin"`
	Time int64            `json:"time"`
	Data []CommissionInfo `json:"data"`
}

func (g GcicsApi) Commission(coinType string, commissionList []CommissionInfo) error {
	c := g.SvcCtx.Config.GcicsConf
	postUrl := c.Host + "/Game/manage/commission"
	timestamp := time.Now().Unix()
	sign, err := util.GenerateSHA1Signature(strconv.FormatInt(timestamp, 10), c.AppId, c.AppSecret)
	if err != nil {
		return err
	}

	// 定义要发送的数据
	data := &CommissionReq{
		Coin: coinType,
		Time: timestamp,
		Data: commissionList,
	}

	// 将数据转换为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 创建请求
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-PD-SIGN", sign)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errorx.NewCodeInternalError(fmt.Sprintf("Failed to withdraw[%d]", resp.StatusCode))
	}

	// 解析JSON响应到BaseResponse结构体
	var postResp BaseResponse
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}

	return g.CheckStatus(postResp.Status, postResp.Code, postResp.Msg)
}

func (g GcicsApi) Logout(coinType string, lambAmount float64) error {
	c := g.SvcCtx.Config.GcicsConf
	postUrl := c.Host + "/Game/api/logout"
	timestamp := time.Now().Unix()
	sign, err := util.GenerateSHA1Signature(strconv.FormatInt(timestamp, 10), c.AppId, c.AppSecret)
	if err != nil {
		return err
	}

	// 创建请求
	req, err := http.NewRequest("POST", postUrl, bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-PD-SIGN", sign)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errorx.NewCodeInternalError(fmt.Sprintf("Failed to withdraw[%d]", resp.StatusCode))
	}

	// 解析JSON响应到BaseResponse结构体
	var postResp BaseResponse
	err = json.Unmarshal(body, &postResp)
	if err != nil {
		return errorx.NewCodeInternalError(fmt.Sprintf("Error unmarshalling JSON: %s", err.Error()))
	}

	return g.CheckStatus(postResp.Status, postResp.Code, postResp.Msg)
}
