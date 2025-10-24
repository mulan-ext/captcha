package captcha_test

import (
	"fmt"
	"testing"

	"github.com/mulan-ext/captcha"
)

func TestCreate(t *testing.T) {
	id, result, _ := captcha.Create()
	fmt.Println("id", id, "result", result)
	if ok, _ := captcha.Check(id, result); !ok {
		t.Fail()
	}
}

func TestCreateBytes(t *testing.T) {
	id, result, _, err := captcha.CreateBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id", id, "result", result)
	if ok, _ := captcha.Check(id, result); !ok {
		t.Fail()
	}
}

func TestCreateB64(t *testing.T) {
	id, result, _, err := captcha.CreateB64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id", id, "result", result)
	if ok, _ := captcha.Check(id, result); !ok {
		t.Fail()
	}
}

func BenchmarkCreate(b *testing.B) {
	for b.Loop() {
		captcha.Create()
	}
}

func BenchmarkCreateBytes(b *testing.B) {
	for b.Loop() {
		captcha.CreateBytes()
	}
}

func BenchmarkCreateB64(b *testing.B) {
	for b.Loop() {
		captcha.CreateB64()
	}
}
