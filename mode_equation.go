package captcha

import (
	"bytes"
	"image"
	"image/color"
	"math"
	"math/rand"
	"strconv"
)

type CaptchaEquation struct {
	*CaptchaBase
	bit int
}

var _ Captcha = (*CaptchaEquation)(nil)

var Equation = DefaultCaptchaEquation()

func DefaultCaptchaEquation() Captcha {
	return NewCaptchaEquation(2)
}

func NewCaptchaEquation(bit int, opts ...Option) Captcha {
	if bit <= 0 {
		panic("bit must be greater than 0")
	}
	std = &CaptchaEquation{
		CaptchaBase: New(
			WithSize(102, 30),
			WithFontByte(fontData, 24),
			WithBackground(color.Black),
			WithFront(color.White),
			WithPoint(100),
			WithLine(6),
		).Options(opts...),
		bit: bit,
	}
	return std
}

func (c *CaptchaEquation) Draw() (image.Image, *CaptchaData) {
	text, result := randomText(c.bit)
	img, cd := c.CaptchaBase.draw(text)
	cd.Result = result
	return img, cd
}

const (
	equationMethodAdd = iota // x + y = z
	equationMethodSub        // x - y = z
	equationMethodMul        // x * y = z
)

func randomText(bit int) (content, result string) {
	var resultNum int64
	buf := new(bytes.Buffer)
	minNum := int64(math.Pow10(bit - 1))
	switch rand.Intn(3) {
	case equationMethodAdd: // left + right = result
		var left, right int64
		resultNum = nBitNum(bit)
		for resultNum < minNum*2 {
			resultNum = nBitNum(bit)
		}
		for left < minNum || right < minNum {
			left = rand.Int63n(resultNum-minNum) + minNum
			right = resultNum - left
		}
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('+')
		buf.WriteString(strconv.FormatInt(resultNum-left, 10))
	case equationMethodSub: // left - right = result
		var left, right int64
		for left < minNum*2 {
			left = nBitNum(bit)
		}
		for resultNum < minNum || right < minNum {
			resultNum = rand.Int63n(left-minNum) + minNum
			right = left - resultNum
		}
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('-')
		buf.WriteString(strconv.FormatInt(right, 10))
	case equationMethodMul: // left * right = result
		left, right := nBitNum(bit), nBitNum(bit)
		resultNum = left * right
		buf.WriteString(strconv.FormatInt(left, 10))
		buf.WriteByte('*')
		buf.WriteString(strconv.FormatInt(right, 10))
	}
	return buf.String(), strconv.FormatInt(resultNum, 10)
}
