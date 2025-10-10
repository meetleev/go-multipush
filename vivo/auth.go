package vivo

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/meetleev/go-multipush/utils"
)

const (
	timeout           = 5
	deviceTokenMax    = 1000
	deviceTokenMin    = 1
	urlBase           = "https://api-push.vivo.com.cn"
	actionAuth        = "message/auth"
	actionSinglePush  = "message/send"
	actionSaveMessage = "message/saveListPayload"
	actionMultiPush   = "message/pushToList"
)

func buildRequest(request *AuthTokenReq) map[string]string {
	request.Timestamp = strconv.FormatInt(time.Now().UTC().UnixNano()/(1e6), 10)

	return map[string]string{
		"appId":     request.AppId,
		"appKey":    request.AppKey,
		"timestamp": request.Timestamp,
		"sign":      generateSign(request),
	}
}

func generateSign(request *AuthTokenReq) string {
	signStr := request.AppId + request.AppKey + request.Timestamp + request.AppSecret
	signStr = strings.Trim(signStr, "")
	return strings.ToLower(utils.MD5([]byte(signStr)))
}

func getUri() string {

	return fmt.Sprintf("%s/%s", urlBase, actionAuth)
}
