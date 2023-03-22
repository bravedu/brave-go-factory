package service

import (
	"encoding/json"
	"fmt"
	"github.com/bravedu/brave-go-factory/config"
	"github.com/bravedu/brave-go-factory/pkg/logger"
	"github.com/bravedu/brave-go-factory/pkg/util"
	"github.com/bravedu/brave-go-factory/pkg/wx"
	"github.com/sirupsen/logrus"
	"github.com/skip2/go-qrcode"
)

func TestAliYunOssUpload() error { //测试上传阿里云OSS,图片生成二维码
	var png []byte
	userShareText := fmt.Sprintf("%s%s%d", "http://www.test.com/", "pay/chat?share_open_id=", 1)
	png, err := qrcode.Encode(userShareText, qrcode.Highest, 50)
	if err != nil {
		logger.AppLog.WithFields(logrus.Fields{
			"err": err,
		}).Info("生成用户二维码失败")
	}
	config.Conf.AliYunCli.OssCli.Cli.UploadByImgByte(png, fmt.Sprintf("%d%s", 1, ".png"), "")
	return nil
}

func TestAliYunSMS() error { //短信发送
	msgCode := util.RandomNumber(6)
	//msgCode = "888888"
	//TODO 阿里云验证码
	config.Conf.AliYunCli.SMSCli.SendSms("19933610888", fmt.Sprintf(`{"code":"%s"}`, msgCode), "SMS_257700803", "签名")
	return nil
}

func TestBaiduAuditCheck(text string) bool { //百度敏感词审核服务
	token, err := config.Conf.BaiDuCli.Audit.GetAccessToken()
	if token == "" || err != nil {
		return false
	}
	res, err := config.Conf.BaiDuCli.Audit.AuditText(token, text)
	if res == false || err != nil {
		return false
	}
	return true
}

type LoginByWechatCodeReq struct {
	Code        string `json:"code" form:"code" binding:"required" msg:"微信code"` //微信CODE
	ShareOpenId int    `json:"share_open_id" form:"share_open_id" msg:"分享人Uid"`  //分享人Uid
}

type LoginByWechatCodeResp struct {
	Token string `json:"token"`
}

func LoginByWechatCodeSvc(req *LoginByWechatCodeReq, resp *LoginByWechatCodeResp) error { //微信登录-根据前端传过来的授权码-扫码登录
	//微信获取openid
	var accUrl = config.Conf.YamlDao.Wxweb.AccessTokenUrl
	wxToken, err := wx.GetWxToken(accUrl + req.Code)
	if err != nil {
		//log.Println("Service--login--wxToken:", wxToken, "---err:", err)
		return err
	}
	//拉取微信用户信息
	wxUserInfo, err := GetWxUserInfo(wxToken.AccessToken, wxToken.Openid)
	if err != nil {
		//log.Println("Service--login--GetWxUserInfo:", wxUserInfo, "---err:", err)
		return err
	}
	//微信用户信息
	fmt.Println(wxUserInfo)
	return nil
}

func LoginWechatH5ByCodeSvc(req LoginByWechatCodeReq, resp LoginByWechatCodeResp) error { //微信H5登录
	//微信获取openid
	var accUrl = config.Conf.YamlDao.Wxh5.AccessTokenUrl
	wxToken, err := wx.GetWxToken(accUrl + req.Code)
	if err != nil {
		return err
	}
	//拉取微信用户信息
	wxUserInfo, err := GetWxUserInfo(wxToken.AccessToken, wxToken.Openid)
	if err != nil {
		return err
	}
	//微信用户信息
	fmt.Println(wxUserInfo)
	return nil
}

type WxUserInfo struct {
	Openid     string `json:"openid"`   //微信OpenId
	UnionId    string `json:"unionid"`  //unionid
	Nickname   string `json:"nickname"` //微信昵称
	Sex        int8   `json:"sex"`      //性别
	Province   string `json:"province"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Headimgurl string `json:"headimgurl"` //头像url
}

func GetWxUserInfo(access_token, openid string) (*WxUserInfo, error) {
	var accUrl = config.Conf.YamlDao.Wxh5.UserinfoUrl
	body, err := wx.CallWxUrl(accUrl + "?access_token=" + access_token + "&openid=" + openid + "&lang=zh_CN")
	if err != nil {
		return nil, err
	}

	var wxUserinfo WxUserInfo
	err = json.Unmarshal(body, &wxUserinfo)
	if err != nil {
		return nil, err
	}
	//logger.AppLog.WithFields(logrus.Fields{
	//	"wx_user_info": wxUserinfo,
	//}).Info("微信用户信息")
	//log.Println("GetWxUserInfo_Unmarshal_wxUserinfo:", wxUserinfo)
	return &wxUserinfo, nil
}
