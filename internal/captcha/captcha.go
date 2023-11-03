package captcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"time"

	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	captchaMap = cmap.New[*CaptchaData]()
)

// Create 创建验证码，返回图片对象image.Image
func Create(s Captcha) (id, result string, img image.Image) {
	img, cd := s.Draw()
	id = uuid.New().String()
	captchaMap.Set(id, cd)
	return id, cd.Result, img
}

// CreateBytes 创建验证码，返回图片数据[]byte
func CreateBytes(s Captcha) (id, result string, buf []byte, err error) {
	id, result, img := Create(s)
	_buf := new(bytes.Buffer)
	err = png.Encode(_buf, img)
	if err != nil {
		return "", "", nil, err
	}
	return id, result, _buf.Bytes(), nil
}

// CreateB64 创建验证码，返回图片base64
func CreateB64(s Captcha) (id, result, data string, err error) {
	id, result, buf, err := CreateBytes(s)
	if err != nil {
		return "", "", "", err
	}
	data = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
	return id, result, data, nil
}

// Check 验证验证码是否正确,返回错误类型
func Check(id, code string) (bool, error) {
	if data, ok := captchaMap.Get(id); ok {
		if data.Expire <= time.Now().Unix() {
			captchaMap.Remove(id)
			return false, ErrCodeExpired
		}
		if data.Result == code {
			captchaMap.Remove(id)
			return true, nil
		}
	}
	return false, ErrCodeInvalid
}

// Check 验证验证码是否正确
func CheckOk(id, code string) bool {
	ok, _ := Check(id, code)
	return ok
}
