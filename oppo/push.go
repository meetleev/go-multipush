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
	error
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
	Data    *authtoken `json:"data"`
}
type authtoken struct {
	Auth_token  string `json:"auth_token"`
	Create_time int    `json:"create_time"`
}

func getHost() string {
	return CN_URL
}

// auth_token
func Auth(app_key string, masterSecret string) (string, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/(1e6), 10)
	params := url.Values{}
	signStr := app_key + timestamp + masterSecret
	params.Set("app_key", app_key)
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
		return "", err
	}
	if 0 == res.Code {
		return res.Data.Auth_token, nil
	}
	return "", HttpError{Code: res.Code, Message: res.Message}
}

func SendSingleMessage(accessToken string, msg *PushSingleMessageReq) error {
	params := url.Values{}
	headers := make(map[string]string)
	bytesData, _ := json.Marshal(msg)
	params.Set("message", string(bytesData))
	params.Set("auth_token", accessToken)
	headers["auth_token"] = accessToken
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(fmt.Sprintf("%s/server/v1/message/notification/unicast", getHost()), []byte(params.Encode()), headers)
	if err != nil {
		return err
	}
	logger.Info("oppo:SendMessage => resp:", string(jsonData))
	return nil
}
