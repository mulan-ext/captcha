package captcha

import (
	"errors"
	"fmt"
)

type CaptchaErr struct {
	ErrCode int
	error
}

func (e CaptchaErr) Error() string {
	return fmt.Sprintf("%d : %s", e.ErrCode, e.error.Error())
}

var (
	ErrCodeExpired = CaptchaErr{ErrCode: 301001, error: errors.New("验证码已过期")}
	ErrCodeInvalid = CaptchaErr{ErrCode: 301002, error: errors.New("验证码无效")}
)
