package util

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var (
	DriverAudio   *base64Captcha.DriverAudio   //音频类型
	DriverString  *base64Captcha.DriverString  //字符串型
	DriverChinese *base64Captcha.DriverChinese //汉字型
	DriverMath    *base64Captcha.DriverMath    //数学型
	DriverDigit   *base64Captcha.DriverDigit   //默认型
)

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha(driver base64Captcha.Driver) (captchaId, captchaB64s string) {
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, captchaB64s, err := c.Generate()
	if err != nil {
		return "", ""
	}
	return captchaId, captchaB64s
}

func VerifyCaptcha(captchaId, verifyValue string) bool {
	if store.Verify(captchaId, verifyValue, true) {
		return true
	}
	return false
}

func DriverMathCnf() *base64Captcha.DriverMath {
	mathType := &base64Captcha.DriverMath{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return mathType
}

// DriverDigitConf 生成图形化数字验证码配置
func DriverDigitConf(h, w, l int) *base64Captcha.DriverDigit {
	digitType := &base64Captcha.DriverDigit{
		Height:   h, //50
		Width:    w, //100
		Length:   l, //5
		MaxSkew:  0.45,
		DotCount: 80,
	}
	return digitType
}

// DriverStringCnf 生成图形化字符串验证码配置
func DriverStringCnf() *base64Captcha.DriverString {
	stringType := &base64Captcha.DriverString{
		Height:          100,
		Width:           50,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSlimeLine,
		Length:          5,
		Source:          "123456789qwertyuiopasdfghjklzxcvb",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return stringType
}

// DriverChineseCnf 生成图形化汉字验证码配置
func DriverChineseCnf() *base64Captcha.DriverChinese {
	chineseType := &base64Captcha.DriverChinese{
		Height:          50,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowSlimeLine,
		Length:          2,
		Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,不想要,的值",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
	return chineseType
}

// DriverAudioCnf 生成图形化数字音频验证码配置
func DriverAudioCnf() *base64Captcha.DriverAudio {
	chineseType := &base64Captcha.DriverAudio{
		Length:   4,
		Language: "zh",
	}
	return chineseType
}
