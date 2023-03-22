package aliyun_mod

import (
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"net/url"
	"strings"
	"sync"
)

var (
	smsPool *SmsPool
	smsOnce sync.Once
)

type SmsPool struct {
	Client *dysmsapi.Client
}

type ALiYunSms struct {
	AccessKey    string `yaml:"access_key"`
	AccessSecret string `yaml:"access_secret"`
	RegionId     string `yaml:"region_id"`
}

func SmsPoolInstance(cnf ALiYunSms) *SmsPool {
	smsOnce.Do(func() {
		smsPool = &SmsPool{
			Client: NewSmsClient(cnf.RegionId, cnf.AccessKey, cnf.AccessSecret),
		}
	})
	return smsPool
}

// SendSmsReply 短信发送响应
type SendSmsReply struct {
	Code    string `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}

func replace(in string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(in))
}

func NewSmsClient(regionId, key, secret string) *dysmsapi.Client {
	config := &sdk.Config{
		EnableAsync:       true,
		GoRoutinePoolSize: 5,
		MaxTaskQueueSize:  1000, //一次请求最大可发送手机号的数量
	}
	credential := credentials.NewAccessKeyCredential(key, secret)
	client, err := dysmsapi.NewClientWithOptions(regionId, config, credential)
	if err != nil {
		panic(err)
	}
	return client
}

func (sp *SmsPool) SendSms(phoneNumbers, templateParam, templateCode, signName string) error {
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phoneNumbers
	request.TemplateCode = templateCode
	request.TemplateParam = templateParam
	request.SignName = signName

	response, err := sp.Client.SendSms(request)
	if err != nil {
		return err
	}
	if response.Code != "OK" {
		return errors.New(response.Message)
	}
	return nil
}
