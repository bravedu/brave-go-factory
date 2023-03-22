package baidu_mod

import (
	"encoding/json"
	"errors"
	"github.com/bravedu/brave-go-factory/pkg/logger"
	"github.com/bravedu/brave-go-factory/pkg/util"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	AccessTokenUrl = "https://aip.baidubce.com/oauth/2.0/token"
	AuditTextUrl   = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"
)

type AuditCnf struct {
	AppId     string `yaml:"app_id"`
	AppKey    string `yaml:"app_key"`
	AppSecret string `yaml:"app_secret"`
}

type AuditClientPool struct {
	Client   *resty.Client
	AuditCnf *AuditCnf
}

var auditClientPool *AuditClientPool
var auditOnce sync.Once

func AuditInstance(cnf *AuditCnf) *AuditClientPool {
	auditOnce.Do(func() {
		auditClientPool = &AuditClientPool{
			Client:   util.GetClient(),
			AuditCnf: cnf,
		}
		//auditClient = util.GetClient()
	})
	return auditClientPool
}

func (a *AuditClientPool) GetAccessToken() (token string, err error) {
	client := a.Client.R()
	//reqBody := map[string]interface{}{
	//	"grant_type":    "client_credentials",
	//	"client_id":     a.AuditCnf.AppId,
	//	"client_secret": a.AuditCnf.AppSecret,
	//}
	url := AccessTokenUrl + "?grant_type=client_credentials&client_id=" + a.AuditCnf.AppKey + "&client_secret=" + a.AuditCnf.AppSecret
	//client.SetBody(reqBody)
	httpRes, err := client.Get(url)
	if err != nil || httpRes.RawResponse.StatusCode != 200 {
		logger.AppLog.WithFields(logrus.Fields{
			"httpRes":     httpRes,
			"RawResponse": httpRes.RawResponse,
		}).Error("Http 获取Token 失败")
		return "", errors.New("获取Token失败")
	}
	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	json.Unmarshal(httpRes.Body(), &result)
	if result.AccessToken == "" {
		return "", errors.New("Token为空")
	}
	//caches.SetBaiDuToken(result.RefreshToken, result.ExpiresIn-100)
	return result.AccessToken, err
}

func (a *AuditClientPool) AuditText(token, text string) (bool, error) {
	//token, _ := caches.GetBaiDuToken()
	//if token == "" {
	//
	//}
	//token, _ := a.GetAccessToken()
	//发起请求
	client := a.Client.R()
	reqBody := map[string]string{
		"text": text,
	}
	client.SetFormData(reqBody)
	client.Header.Set("Accept", "application/x-www-form-urlencoded")
	httpRes, err := client.Post(AuditTextUrl + "?access_token=" + token)
	if err != nil || httpRes.RawResponse.StatusCode != 200 {
		logger.AppLog.WithFields(logrus.Fields{
			"httpRes":     httpRes,
			"RawResponse": httpRes.RawResponse,
		}).Error("Http 获取Token 失败")
		return false, errors.New("获取Token失败")
	}
	var result struct {
		Conclusion string `json:"conclusion"`
	}
	json.Unmarshal(httpRes.Body(), &result)
	if result.Conclusion != "合规" {
		return false, err
	}
	return true, nil
}
