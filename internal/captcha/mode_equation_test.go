package captcha

import (
	"fmt"
	"testing"

	"github.com/virzz/logger"
)

// func TestCheckEquation(t *testing.T) {
// 	c := NewCaptchaEquation()
// }

func TestEquationDraw(t *testing.T) {
	c := NewCaptchaEquation()
	_, d := c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
	_, d = c.Draw()
	fmt.Printf("%+v\n", d)
}

func TestEquationBound(t *testing.T) {
	fontFace, err := GetFontFace(fontData, 24.0)
	if err != nil {
		t.Error(err)
	}
	var (
		c          string
		w, h       int
		maxW, maxH int
	)
	for i := 0; i < 100; i++ {
		c, _ = randomText()
		w, h = BoundString(c, fontFace)
		if w > maxW {
			maxW = w
		}
		if h > maxH {
			maxH = h
		}
	}
	logger.SuccessF("maxW: %d maxH: %d", maxW, maxH)
}
