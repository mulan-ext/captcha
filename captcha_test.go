package captcha

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestEquation(t *testing.T) {
	SetCaptcha(Equation)
	// Create()
	img, cd := std.Draw()
	fmt.Printf("%+v", cd)
	// 输出文件
	f, err := os.OpenFile("./test.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		t.Error(err)
	}
	exec.Command("open", "./test.png").Run()
}

func TestCreateB64(t *testing.T) {
	id, result, _, err := CreateB64()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
	fmt.Println(result)
	fmt.Println(time.Now().Unix())
	if !CheckOk(id, result) {
		t.Fail()
	}
}

func BenchmarkCreate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Create()
	}
}
