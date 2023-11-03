package equation

import (
	"fmt"
	"os"
	"testing"

	"github.com/virzz/logger"
)

func TestCreateB64(t *testing.T) {
	id, result, _, err := CreateB64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
	fmt.Println(result)
	if CheckOk(id, "xxxxx") {
		t.Fail()
	}
	if !CheckOk(id, result) {
		t.Fail()
	}
}

func TestEquation(t *testing.T) {
	id, res, buf, err := CreateBytes()
	if err != nil {
		t.Error(err)
	}
	// 输出文件
	f, err := os.OpenFile("./test.png", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	_, err = f.Write(buf)
	if err != nil {
		t.Error(err)
	}
	logger.InfoF("id: %s res: %s", id, res)
	// exec.Command("open", "./test.png").Run()
}

func TestCreate(t *testing.T) {
	for i := 0; i < 100; i++ {
		Create()
	}
}
