package captcha

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"sync"
	"time"

	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map/v2"
)

var (
	captchaMap = cmap.New[*CaptchaData]()
	once       sync.Once
	debug      bool = false
	std        Captcha
)

func init() {
	std = Equation
}

func init() {
	go once.Do(func() {
		for {
			time.Sleep(time.Minute)
			captchaMap.IterCb(func(key string, value *CaptchaData) {
				if value.Expire <= time.Now().Unix() {
					captchaMap.Remove(key)
				}
			})
		}
	})
}

func SetCaptcha(m Captcha) { std = m }
func Debug()               { debug = !debug }

// Create 创建验证码，返回图片对象image.Image
func Create() (id string, result string, img image.Image) {
	return create(std)
}

// CreateBytes 创建验证码，返回图片数据[]byte
func CreateBytes() (id string, result string, buf []byte, err error) {
	return createBytes(std)
}

// CreateB64 创建验证码，返回图片base64
func CreateB64() (id string, result string, data string, err error) {
	return createB64(std)
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

func create(s Captcha) (id, result string, img image.Image) {
	img, cd := s.Draw()
	id = uuid.New().String()
	captchaMap.Set(id, cd)
	return id, cd.Result, img
}

func createBytes(s Captcha) (id, result string, buf []byte, err error) {
	id, result, img := create(s)
	_buf := new(bytes.Buffer)
	err = png.Encode(_buf, img)
	if err != nil {
		return "", "", nil, err
	}
	return id, result, _buf.Bytes(), nil
}

func createB64(s Captcha) (id, result, data string, err error) {
	id, result, buf, err := createBytes(s)
	if err != nil {
		return "", "", "", err
	}
	data = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf)
	return id, result, data, nil
}
