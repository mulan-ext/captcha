package captcha

import (
	crand "crypto/rand"
	"encoding/hex"
	"image"
	"image/color"
	"math"
	"math/rand"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

const (
	// 0x30-0x39 1234567890
	// 0x41-0x5a ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// 0x61-0x7a abcdefghijklmnopqrstuvwxyz
	letterBytes   = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandomCode() string {
	return strconv.Itoa(rand.Intn(900000) + 100000)
}

func RandomBytes(n int) []byte {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return b
}

func RandomBytesHex(n int) string {
	return hex.EncodeToString(RandomBytes(n))
}

func RandomString(n int) string {
	return string(RandomBytes(n))
}

func sign(x int) int {
	if x > 0 {
		return 1
	}
	return -1
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func drawRect(fr fixed.Rectangle26_6) (x, y, width, height int) {
	return 0, 0, int(fr.Max.X+fr.Min.X) >> 6, int(fr.Max.Y+fr.Min.Y)>>6 + 1
}

func unFixedRect(b fixed.Rectangle26_6) image.Rectangle {
	return image.Rect(drawRect(b))
}

func getFontFace(buf []byte, fontSize float64, dpi ...float64) (font.Face, error) {
	obj, err := sfnt.Parse(buf)
	if err != nil {
		return nil, err
	}
	_dpi := 72.0
	if len(dpi) > 0 {
		_dpi = dpi[0]
	}
	fontFace, err := opentype.NewFace(obj, &opentype.FaceOptions{
		Size: fontSize, DPI: _dpi, Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, err
	}
	return fontFace, nil
}

func randColor(alpha ...bool) color.Color {
	buf := make([]byte, 4)
	crand.Read(buf)
	a := uint8(0xff)
	if len(alpha) > 0 && alpha[0] {
		a = buf[3]
	}
	return color.RGBA{uint8(buf[0]), uint8(buf[1]), uint8(buf[2]), a}
}

func nBitNum(bit int) int64 {
	x := math.Pow10(bit - 1)
	return rand.Int63n(int64(math.Pow10(bit)-x)) + int64(x)
}
