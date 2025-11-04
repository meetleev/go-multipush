package oppo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/meetleev/go-multipush/utils"
	logger "github.com/sirupsen/logrus"
)

const (
	CN_URL     = "https://api.push.oppomobile.com"
	GlOBAL_URL = "https://api-intl.push.oppomobile.com"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (o HttpError) Error() string {
	return o.Message
}

type AuthResp struct {
	// https://open.oppomobile.com/documentation/page/info?id=11235
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    *AuthToken `json:"data"`
}
type AuthToken struct {
	AuthToken string `json:"auth_token"`
	CreatedAt int    `json:"create_time"`
}

func getHost() string {
	return CN_URL
}

type OppoPushClient struct {
	appPushKey   string
	masterSecret string

	tokenData *AuthToken
}

func NewPushClient(appPushKey, masterSecret string) *OppoPushClient {
	return &OppoPushClient{appPushKey: appPushKey, masterSecret: masterSecret}
}

// auth_token
func (c *OppoPushClient) Auth() (*AuthToken, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/(1e6), 10)
	params := url.Values{}
	signStr := c.appPushKey + timestamp + c.masterSecret
	params.Set("app_key", c.appPushKey)
	params.Set("sign", utils.SHA256([]byte(signStr)))
	params.Set("timestamp", timestamp)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(fmt.Sprintf("%s/server/v1/auth", getHost()), []byte(params.Encode()), headers)
	if err != nil {
		logger.Error("Oppo:Auth", err)
	}
	res := AuthResp{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		return nil, err
	}
	if 0 == res.Code {
		return res.Data, nil
	}
	return nil, HttpError{Code: res.Code, Message: res.Message}
}

func (c *OppoPushClient) SendSingleMessage(msg *PushSingleMessageReq) (*PushMessageResp, error) {
	now := time.Now().UnixMilli()
	if nil == c.tokenData {
		tokenData, err := c.Auth()
		if err != nil {
			return nil, err
		}
		c.tokenData = tokenData
	} else {
		if 60*60000 > now-int64(c.tokenData.CreatedAt) {
			c.tokenData = nil
			return c.SendSingleMessage(msg)
		}
	}
	accessToken := c.tokenData.AuthToken
	params := url.Values{}
	headers := make(map[string]string)
	bytesData, _ := json.Marshal(msg)
	params.Set("message", string(bytesData))
	// params.Set("auth_token", accessToken)
	headers["auth_token"] = accessToken
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(fmt.Sprintf("%s/server/v1/message/notification/unicast", getHost()), []byte(params.Encode()), headers)
	if err != nil {
		return nil, err
	}
	resp := &PushMessageResp{}
	err = json.Unmarshal(jsonData, resp)
	logger.Info("oppo:SendMessage => resp:", string(jsonData))
	return resp, err
}
