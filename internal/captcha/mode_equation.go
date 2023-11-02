package captcha

import (
	"bytes"
	"image"
	"image/color"
	"math/rand"
	"strconv"
)

type CaptchaEquation struct {
	*CaptchaBase
}

var _ Captcha = (*CaptchaEquation)(nil)

func DefaultCaptchaEquation() *CaptchaEquation {
	return &CaptchaEquation{
		CaptchaBase: New(
			WithSize(102, 30), WithFontByte(fontData, 24),
			WithBackground(color.Black), WithFront(color.White),
			WithPoint(100),
			WithLine(10),
		),
	}
}

func NewCaptchaEquation(opts ...Option) *CaptchaEquation {
	return &CaptchaEquation{
		CaptchaBase: New(
			WithSize(102, 30), WithFontByte(fontData, 24),
			WithBackground(color.Black), WithFront(color.White),
			WithPoint(100),
			WithLine(10),
		).Options(opts...),
	}
}

const (
	equationMethodAdd = iota + 1 // x + y = z
	equationMethodSub            // x - y = z
	equationMethodMul            // x * y = z
)

func randomText() (content, result string) {
	var resultNum int64
	buf := new(bytes.Buffer)
	switch rand.Intn(3) + 1 {
	case equationMethodAdd: // left + right = result
		resultNum = rand.Int63n(990) + 10
		left := rand.Int63n(resultNum)
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('+')
		buf.WriteString(strconv.FormatInt(resultNum-left, 10))
	case equationMethodSub: // left - right = result
		resultNum = rand.Int63n(990) + 10
		left := rand.Int63n(resultNum)
		right := resultNum - left
		if right > left {
			left, right = right, left
		}
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('-')
		buf.WriteString(strconv.FormatInt(right, 10))
	case equationMethodMul: // left * right = result
		left, right := rand.Int63n(90)+10, rand.Int63n(90)+10
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('*')
		buf.WriteString(strconv.FormatInt(right, 10))
		resultNum = left * right
	}
	return buf.String(), strconv.FormatInt(resultNum, 10)
}

func (c *CaptchaEquation) Draw() (image.Image, *CaptchaData) {
	text, result := randomText()
	img, cd := c.CaptchaBase.draw(text)
	cd.Result = result
	return img, cd
}
