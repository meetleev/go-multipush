package mi

import (
	"encoding/json"
	"strings"

	logger "github.com/sirupsen/logrus"

	"github.com/meetleev/go-multipush/utils"
)

var (
	v2URI      string = "https://api.xmpush.xiaomi.com/v2"
	v3URI      string = "https://api.xmpush.xiaomi.com/v3"
	defaultURI string = v3URI
)

type MiPush struct {
	AppID     string
	AppKey    string
	AppSecret string
}

// 推送消息给指定的一些 regid
func (miPush *MiPush) SendMultAlias(message PushMessage, regIDs []string) (*PushMessageResp, error) {
	req := PushMessageReq{
		PushMessage: message,
		Alias:       strings.Join(regIDs, ","),
	}
	return miPush.SendAlias(&req)
}

// 推送消息给指定的一些 regid
func (miPush *MiPush) SendMultRegId(message PushMessage, regIDs []string) (*PushMessageResp, error) {
	req := PushMessageReq{
		PushMessage:    message,
		RegistrationId: strings.Join(regIDs, ","),
	}
	return miPush.SendRegID(&req)
}

// message/alias
func (miPush *MiPush) SendAlias(req *PushMessageReq) (*PushMessageResp, error) {
	bytesData, _ := json.Marshal(req)
	uri := defaultURI + "/message/alias"
	headers := make(map[string]string)
	headers["Authorization"] = "key=" + miPush.AppSecret
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(uri, bytesData, headers)
	if err != nil {
		logger.Errorf("httpPost(%v, %v, %v) error(%v)", uri, string(bytesData), headers, err)
		return nil, err
	}
	res := PushMessageResp{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		logger.Errorf("json.Unmarshal(%s,&res) error(%v)", jsonData, err)
		return nil, err
	}
	logger.Debugf("result(%s)", jsonData)
	if res.Code != 0 {
		// 21301:认证失败
		// 21305:缺失必选参数
		// 10016:缺失必选参数
		logger.Errorf("result(%s)", jsonData)
		return nil, HttpError{Code: res.Code, Message: res.Description}
	}

	return &res, nil
}

//	params.Set("channel_id", "101351")
//
// 推送消息给指定的regid
func (miPush *MiPush) SendRegID(req *PushMessageReq) (*PushMessageResp, error) {
	bytesData, _ := json.Marshal(req)
	uri := defaultURI + "/message/regid"
	headers := make(map[string]string)
	headers["Authorization"] = "key=" + miPush.AppSecret
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	jsonData, err := utils.HttpPost(uri, bytesData, headers)
	if err != nil {
		logger.Errorf("httpPost(%v, %v, %v) error(%v)", uri, string(bytesData), headers, err)
		return nil, err
	}

	logger.Debugf(" SendRegID(%s)", string(bytesData))

	res := PushMessageResp{}
	if err := json.Unmarshal(jsonData, &res); err != nil {
		logger.Errorf("json.Unmarshal(%s,&res) error(%v)", jsonData, err)
		return nil, err
	}
	logger.Debugf("result(%s)", jsonData)

	if res.Code != 0 {
		// 21301:认证失败
		// 21305:缺失必选参数
		// 10016:缺失必选参数
		logger.Errorf("result(%s)", jsonData)
		return nil, HttpError{Code: res.Code, Message: res.Description}
	}
	return &res, nil
}
