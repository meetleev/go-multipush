package vivo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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

func Auth(appId string, appKey string, appSecret string) (string, error) {
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
	logger.Debugf("vivo:Auth=>resp:[%s]", res)
	authStructs.AuthStructs = resp
	authStructs.Date = t
	return resp.AuthToken, err
}

func SendMessage(token string, title string, content string, regId string) error {
	song := make(map[string]string)
	song["regId"] = regId
	song["notifyType"] = "1"
	song["title"] = title     //"标题"
	song["content"] = content // "内容"
	song["timeToLive"] = "86400"
	song["skipType"] = "1"
	song["skipContent"] = "1"
	song["networkType"] = "-1"
	song["pushMode"] = "0"
	song["classification"] = "1"
	song["requestId"] = regId
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json;charset=utf-8"
	headers["authToken"] = token
	bytesData, _ := json.Marshal(song)
	res, err := httpPost(URL+"/message/send", bytesData, headers)
	if err != nil {
		return err
	}
	logger.Debugf("vivo:MessageSend=>token:[%s] resp:[%s]", token, res)
	return nil
}

func httpPost(url string, msg []byte, headers map[string]string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewReader(msg))
	if err != nil {
		return "", err
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}
