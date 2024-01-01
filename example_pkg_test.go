package captcha_test

import (
	"fmt"

	"github.com/virzz/captcha"
)

func ExampleCreateB64() {
	id, result, data, err := captcha.CreateB64()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)
	fmt.Println(data)
	if captcha.CheckOk(id, "xxxx") {
		fmt.Println("test check")
	} else if captcha.CheckOk(id, result) {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}
}
