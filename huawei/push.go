package huawei

import (
	"crypto/rsa"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/meetleev/go-multipush/utils"
)

const V3_URL = "https://push-api.cloud.huawei.com/v3"

type ServiceAccountKey struct {
	KeyID      string `json:"key_id"`
	SubAccount string `json:"sub_account"`
	PrivateKey string `json:"private_key"`
	ProjectId  string `json:"project_id"`
}

type PushClient struct {
	keyID      string
	subAccount string
	projectId  string
	privateKey *rsa.PrivateKey
}

func NewPushClient(keyFilePath string) (*PushClient, error) {
	saKey, err := loadServiceAccountKey(keyFilePath)
	if nil != err {
		return nil, err
	}
	formattedPrivateKey, err := formatPrivateKey(saKey.PrivateKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(formattedPrivateKey))
	if err != nil {
		return nil, err
	}
	return &PushClient{privateKey: privateKey, keyID: saKey.KeyID,
		subAccount: saKey.SubAccount, projectId: saKey.ProjectId}, nil
}

func (c *PushClient) Auth() (string, error) {
	token, err := buildJWTToken(c.keyID, c.subAccount)
	if err != nil {
		return "", err
	}
	return token.SignedString(c.privateKey)
}

func (c *PushClient) SendMessage(msg PushMessage) (*PushResp, error) {
	token, err := c.Auth()
	if nil != err {
		return nil, err
	}
	url := fmt.Sprintf("%s/%s/messages:send", V3_URL, c.projectId)
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", token)
	var pushType PushType
	_, ok := msg.Payload.(*AlertPayload)
	if ok {
		pushType = PushTypeAlert
	}
	byteData, _ := json.Marshal(msg)
	headers["push-type"] = fmt.Sprintf("%d", pushType)
	jsonData, err := utils.HttpPost(url, byteData, headers)
	if nil != err {
		return nil, err
	}
	resp := &PushResp{}
	err = json.Unmarshal(jsonData, resp)
	return resp, err
}

// buildJWTToken 构造 JWT token 对象
func buildJWTToken(keyID, subAccount string) (*jwt.Token, error) {
	now := time.Now().UTC()
	iat := now.Unix()
	exp := iat + 3600 // token 过期时间：一小时后

	claims := jwt.MapClaims{
		// 实际开发时请将公网地址存储在配置文件或数据库
		"aud": "https://oauth-login.cloud.huawei.com/oauth2/v3/token",
		"iss": subAccount,
		"exp": exp,
		"iat": iat,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)

	// 设置 header
	token.Header["kid"] = keyID
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "PS256"

	return token, nil
}

// loadServiceAccountKey 从 JSON 文件加载服务账号密钥
func loadServiceAccountKey(filename string) (*ServiceAccountKey, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %w", err)
	}

	var saKey ServiceAccountKey
	if err := json.Unmarshal(data, &saKey); err != nil {
		return nil, fmt.Errorf("failed to parse key file: %w", err)
	}

	if saKey.KeyID == "" || saKey.SubAccount == "" || saKey.PrivateKey == "" {
		return nil, errors.New("invalid service account key file: missing required fields")
	}

	return &saKey, nil
}

// formatPrivateKey 格式化私钥字符串为 PEM 格式
func formatPrivateKey(privateKeyStr string) (string, error) {
	trimmed := strings.TrimSpace(privateKeyStr)

	// 如果已经是 PEM 格式，则直接返回
	if strings.HasPrefix(trimmed, "-----BEGIN PRIVATE KEY-----") &&
		strings.HasSuffix(trimmed, "-----END PRIVATE KEY-----") {
		return trimmed, nil
	}

	block, _ := pem.Decode([]byte(trimmed))
	if block == nil {
		return "", errors.New("failed to decode PEM block")
	}

	pemBytes := pem.EncodeToMemory(block)
	if pemBytes == nil {
		return "", errors.New("failed to encode private key to PEM format")
	}

	return string(pemBytes), nil
}
