package captcha

import (
	_ "embed"
	"image"
	"image/color"
	"time"

	"github.com/virzz/logger"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

type Captcha interface {
	Draw() (image.Image, *CaptchaData) // 验证码生成逻辑并绘制图片
}

type CaptchaBase struct {
	fontFace          font.Face
	width, height     int
	background, front color.Color
	fontSize          float64
	expire            int64
	line, point       int
}

//go:embed monaco_ascii.ttf
var fontData []byte

var _ Captcha = (*CaptchaBase)(nil)

func New(opts ...Option) *CaptchaBase {
	// Default
	c := &CaptchaBase{
		fontSize:   14,
		width:      80,
		height:     30,
		expire:     300,
		background: color.White,
		front:      color.Black,
	}
	// Option
	for _, opt := range opts {
		opt(c)
	}
	// Default Font
	if c.fontFace == nil {
		obj, _ := sfnt.Parse(fontData)
		fontFace, err := opentype.NewFace(obj, &opentype.FaceOptions{
			Size: c.fontSize, DPI: 72, Hinting: font.HintingNone,
		})
		if err != nil {
			logger.Error(err)
		} else {
			c.fontFace = fontFace
		}
	}
	return c
}

type CaptchaData struct {
	Content string
	Result  string
	Expire  int64
}

func (c *CaptchaBase) Options(opts ...Option) *CaptchaBase {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *CaptchaBase) Draw() (image.Image, *CaptchaData) {
	panic("No implementation found")
}

func (c *CaptchaBase) draw(data string) (image.Image, *CaptchaData) {
	img := newImage(c.width, c.height)
	img.fillBkg(c.background)               // 背景
	img.drawText(data, c.fontFace, c.front) // 内容
	// 干扰点
	img.drawPoints(c.point)
	// 干扰线
	img.drawLines(c.line)
	// TODO: 旋转
	// img.rotate()
	// TODO: 扭曲
	// img.distort()
	return img, &CaptchaData{Content: data, Expire: c.expire + time.Now().Unix()}
}
