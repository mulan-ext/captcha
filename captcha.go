package captcha

import (
	"image"

	"github.com/virzz/captcha/internal/captcha"
)

var (
	debug bool = false
	std   captcha.Captcha
	mode  = captcha.ModeEquation
)

func init() {
	SetMode(mode)
}

func SetMode(m captcha.ModeType) {
	mode = m
	switch m {
	case captcha.ModeEquation:
		std = captcha.DefaultCaptchaEquation()
	}
}

func Debug() {
	debug = !debug
}

func Check(id string, code string) (bool, error) {
	return captcha.Check(id, code)
}

func CheckOk(id string, code string) bool {
	return captcha.CheckOk(id, code)
}

func Create() (id string, result string, img image.Image) {
	return captcha.Create(std)
}

func CreateBytes() (id string, result string, buf []byte, err error) {
	return captcha.CreateBytes(std)
}

func CreateB64() (id string, result string, data string, err error) {
	return captcha.CreateB64(std)
}
