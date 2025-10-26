package honor

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/meetleev/go-multipush/utils"
	logger "github.com/sirupsen/logrus"
)

const (
	BASE_URL = "https://iam.developer.honor.com"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (o HttpError) Error() string {
	return o.Message
}

type AuthResp struct {
	// 应用级Access Token
	AccessToken string `json:"access_token"`
	// Access Token的剩余有效期，单位：秒。
	ExpiredAt int64 `json:"expires_in"`
	// 固定返回Bearer
	TokenType string `json:"token_type"`
}

type HonorPushClient struct {
	ClientId     string
	ClientSecret string
	AppId        string

	tokenData *AuthResp
}

// auth_token
func (c *HonorPushClient) Auth() (*AuthResp, error) {
	params := url.Values{}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", c.ClientId)
	params.Set("client_secret", c.ClientSecret)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(fmt.Sprintf("%s/auth/token", BASE_URL), []byte(params.Encode()), headers)
	if err != nil {
		logger.Error("Honor:Auth", err)
	}
	res := AuthResp{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *HonorPushClient) SendMessage(msg *PushMessage) (*PushMessageResp, error) {
	now := time.Now().UnixMilli()
	if nil == c.tokenData {
		tokenData, err := c.Auth()
		if err != nil {
			return nil, err
		}
		c.tokenData = tokenData
	} else {
		if 60000 > now-c.tokenData.ExpiredAt {
			c.tokenData = nil
			return c.SendMessage(msg)
		}
	}
	headers := make(map[string]string)
	bytesData, _ := json.Marshal(msg)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", c.tokenData.AccessToken)
	headers["timestamp"] = fmt.Sprintf("%d", now)
	headers["Content-Type"] = "application/json"
	jsonData, err := utils.HttpPost(fmt.Sprintf("%s/api/v1/%s/sendMessage", BASE_URL, c.AppId), bytesData, headers)
	if err != nil {
		return nil, err
	}
	logger.Info("Honor:SendMessage => resp:", string(jsonData))
	resp := &PushMessageResp{}
	err = json.Unmarshal(jsonData, resp)
	return resp, err
}
