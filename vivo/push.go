package vivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/meetleev/go-multipush/utils"

	logger "github.com/sirupsen/logrus"
)

const URL = "https://api-push.vivo.com.cn"

type authStruct struct {
	Result    int    `json:"result"`
	AuthToken string `json:"authToken"`
	Desc      string `json:"desc"`
}

type authStructDate struct {
	Date        time.Time `json:"date"`
	AuthStructs authStruct
}

var authStructs = authStructDate{}

func Auth(appId int, appKey string, appSecret string) (string, error) {
	t := time.Now()
	var authtoken = buildRequest(&AuthTokenReq{
		AppId:     appId,
		AppKey:    appKey,
		AppSecret: appSecret,
	})
	fmt.Println(authtoken)
	bytesData, err := json.Marshal(authtoken)
	if err != nil {
		return "", err
	}
	res, err := http.Post(URL+"/message/auth",
		"application/json;charset=utf-8", bytes.NewBuffer(bytesData))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	resp := authStruct{}

	if err := json.Unmarshal(content, &resp); err != nil {
		return "", err
	}
	logger.Debugf("vivo:Auth=>resp:[%s]", content)
	authStructs.AuthStructs = resp
	authStructs.Date = t
	return resp.AuthToken, err
}

func SendSingleMessage(accessToken string, req *PushSingleMessageReq) error {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["authToken"] = accessToken
	bytesData, _ := json.Marshal(req)
	res, err := utils.HttpPost(URL+"/message/send", bytesData, headers)
	if err != nil {
		return err
	}
	logger.Debugf("vivo:MessageSend=>accessToken:[%s] resp:[%s]", accessToken, string(res))
	return nil
}
