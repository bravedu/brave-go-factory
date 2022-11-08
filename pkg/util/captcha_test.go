package util

import "testing"

func TestGenerateCaptcha(t *testing.T) {
	captchaId, captchaB64s := GenerateCaptcha(DriverDigitConf(50, 100, 5))
	t.Log("captchaId : ", captchaId)
	t.Log("captchaB64s : ", captchaB64s)
}

func TestVerifyCaptcha(t *testing.T) {
	captchaId := ""
	validateVal := ""
	validateCodeRes := VerifyCaptcha(captchaId, validateVal)
	t.Log(validateCodeRes)
}
