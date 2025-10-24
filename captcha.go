package captcha

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/mulan-ext/captcha/core"
	"github.com/mulan-ext/captcha/equation"
)

var std core.ICaptcha = equation.NewEquation(2)

func SetGlobal(m core.ICaptcha) { std = m }

// Create 创建验证码，返回图片对象image.Image
func Create() (id, result string, img image.Image) {
	id, cd, img := std.Draw()
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, r)
		return std.Create()
	}
	return id, cd.Result, img
}

// CreateBytes 创建验证码，返回图片数据[]byte
func CreateBytes() (id, result string, buf []byte, err error) {
	id, result, img := std.Create()
	_buf := new(bytes.Buffer)
	err = png.Encode(_buf, img)
	if err != nil {
		return "", "", nil, err
	}
	return id, result, _buf.Bytes(), nil
}

// CreateB64 创建验证码，返回图片base64
func CreateB64() (id, result, data string, err error) {
	id, result, buf, err := std.CreateBytes()
	if err != nil {
		return "", "", "", err
	}
	data = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
	return id, result, data, nil
}

func Check(id, code string) (bool, error) {
	return std.Check(id, code)
}
