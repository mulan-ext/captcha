package captcha

import (
	"fmt"
	"image/png"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/virzz/captcha/internal/captcha"
	"github.com/virzz/logger"
)

func TestEquation(t *testing.T) {
	SetMode(captcha.ModeEquation)
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
	logger.Info(id)
	logger.Info(result)
	logger.Info(time.Now().Unix())
	if !CheckOk(id, result) {
		t.Fail()
	}

}
