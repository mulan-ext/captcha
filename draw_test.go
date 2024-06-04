package captcha

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"os/exec"
	"testing"

	"github.com/virzz/logger"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func TestSign(t *testing.T) {
	for i := -100; i < 100; i++ {
		fmt.Println(i, sign(i))
	}
}

// 字体，字体大小，字符串长度 ~ 图片尺寸
// 字符串区域检测
func TestCheck(t *testing.T) {
	var (
		fontSize = 28.0
		content  = "88*99"
	)
	obj, _ := sfnt.Parse(fontData)
	fontFace, _ := opentype.NewFace(obj, &opentype.FaceOptions{
		Size: fontSize, DPI: 72, Hinting: font.HintingFull,
	})
	fmt.Println("文本长度: ", len(content))
	m := fontFace.Metrics()
	fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", int(m.Ascent>>6))
	fmt.Printf("%+v\n", int(m.Descent>>6))
	fmt.Printf("%+v\n", int(m.Height>>6))
	fmt.Printf("%+v\n", int(m.XHeight>>6))
	fmt.Printf("%+v\n", int(m.CapHeight>>6))
	fmt.Println("字体大小: ", int(m.Ascent>>6))
	// 文本画布
	drawer := &font.Drawer{Src: image.Black, Face: fontFace, Dot: fixed.P(0, int(fontSize))}
	db, _ := drawer.BoundString(content)
	fmt.Printf("%+v\n", db)
	var _, _, width, height = drawRect(db)
	// 绘制文本
	img := newImage(width, height).fillBkg(color.White)
	fmt.Println("图片大小: ", img.Bounds())
	drawer.Dst = img
	drawer.DrawString(content)
	// 输出文件
	f, err := os.OpenFile("./test.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		logger.Error(err)
		t.Fail()
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		logger.Error(err)
		t.Fail()
	}
	exec.Command("open", "./test.png").Run()
	defer os.Remove("./test.png")
}

func TestDrawArcLine2(t *testing.T) {
	img := newImage(100, 40).fillBkg(color.White)
	img.drawArcLine(20, 20, 5, 8, 23, 43, randColor(true))
}
