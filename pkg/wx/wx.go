package wx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bravedu/brave-go-factory/pkg/http"
	"io/ioutil"
)

type WxSession struct {
	SessionKey string `json:"-"`
	OpenId     string `json:"openid"`
}

type WxAccessToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

func CallWxUrl(url string) ([]byte, error) {
	res, err := http.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	//logger.AppLog.WithFields(logrus.Fields{
	//	"wx_data_uinfo": data,
	//}).Info("微信用户解码信息")
	if _, ok := data["errcode"]; ok {
		//为了兼容下面错误和结果一起返回的格式
		// {
		//  "errcode":0,
		//  "errmsg":"ok",
		//  "ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
		//  "expires_in":7200
		//}
		if data["errmsg"] == "ok" {
			return body, nil
		}
		return nil, errors.New(fmt.Sprintf("%s", data["errmsg"]))
	}
	return body, nil
}

func GetWxToken(url string) (*WxAccessToken, error) {
	body, err := CallWxUrl(url)
	if err != nil {
		return nil, err
	}
	// wx返回的是加密access_token，需验证签名并解密
	//TODO

	var wxtoken WxAccessToken
	err = json.Unmarshal(body, &wxtoken)
	if err != nil {
		return nil, err
	}
	return &wxtoken, nil
}
