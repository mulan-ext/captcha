package captcha

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/virzz/logger"
)

func TestEquationDraw(t *testing.T) {
	c := NewCaptchaEquation(2)
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
	fontFace, err := getFontFace(fontData, 24.0)
	if err != nil {
		t.Error(err)
	}
	var (
		c          string
		w, h       int
		maxW, maxH int
	)
	for i := 0; i < 100; i++ {
		c, _ = randomText(2)
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

func TestEquationMethodAdd(t *testing.T) {
	bit := 2
	minNum := int64(math.Pow10(bit - 1))
	// left + right = result
	for i := 0; i < 20; i++ {
		resultNum := nBitNum(bit)
		for resultNum < minNum*2 {
			resultNum = nBitNum(bit)
		}
		var left, right int64
		for left < minNum || right < minNum {
			left = rand.Int63n(resultNum-minNum) + minNum
			right = resultNum - left
		}
		logger.InfoF("%d + %d = %d", left, right, resultNum)
	}
}

func TestEquationMethodSub(t *testing.T) {
	bit := 2
	minNum := int64(math.Pow10(bit - 1))
	var resultNum int64
	// left - right = result
	for i := 0; i < 20; i++ {
		var left, right int64
		for left < minNum*2 {
			left = nBitNum(bit)
		}
		for resultNum < minNum || right < minNum {
			resultNum = rand.Int63n(left-minNum) + minNum
			right = left - resultNum
		}
		logger.InfoF("%d - %d = %d", left, right, resultNum)
	}
}

func TestEquationMethodMul(t *testing.T) {
	bit := 2
	var resultNum int64
	// left * right = result
	for i := 0; i < 20; i++ {
		var left, right = nBitNum(bit), nBitNum(bit)
		resultNum = left * right
		logger.InfoF("%d * %d = %d", left, right, resultNum)
	}
}
