package captcha

import "errors"

type CaptchaErr struct {
	ErrCode int
	error
}

var (
	ErrCodeExpired = CaptchaErr{ErrCode: 301001, error: errors.New("验证码已过期")}
	ErrCodeInvalid = CaptchaErr{ErrCode: 301002, error: errors.New("验证码无效")}
)
