package captcha

import (
	"image/color"
	"os"

	"github.com/virzz/logger"
	"golang.org/x/image/font"
)

type Option func(cb *CaptchaBase)

func WithSize(width, hight int) Option {
	return func(cb *CaptchaBase) { cb.height, cb.width = hight, width }
}
func WithBackground(color color.Color) Option {
	return func(cb *CaptchaBase) { cb.background = color }
}
func WithFront(color color.Color) Option {
	return func(cb *CaptchaBase) { cb.front = color }
}
func WithExpire(t int64) Option {
	return func(cb *CaptchaBase) { cb.expire = t }
}
func WithLine(n int) Option {
	return func(cb *CaptchaBase) { cb.line = n }
}
func WithPoint(n int) Option {
	return func(cb *CaptchaBase) { cb.point = n }
}

func WithFontSize(size float64) Option {
	return func(cb *CaptchaBase) { cb.fontSize = size }
}
func WithFont(fontFace font.Face, size float64) Option {
	return func(cb *CaptchaBase) { cb.fontFace, cb.fontSize = fontFace, size }
}
func WithFontName(fontName string, fontSize float64, dpi ...float64) Option {
	return func(cb *CaptchaBase) {
		buf, err := os.ReadFile(fontName)
		if err != nil {
			logger.Error(err)
		} else {
			WithFontByte(buf, fontSize, dpi...)(cb)
		}
	}
}
func WithFontByte(buf []byte, fontSize float64, dpi ...float64) Option {
	return func(cb *CaptchaBase) {
		fontFace, err := GetFontFace(buf, fontSize, dpi...)
		if err != nil {
			logger.Error(err)
			return
		}
		WithFont(fontFace, fontSize)(cb)
	}
}
