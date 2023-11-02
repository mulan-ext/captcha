package equation

import "fmt"

func ExampleCreateB64() {
	id, result, data, err := CreateB64()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(id)
	fmt.Println(data)

	if CheckOk(id, "xxxx") {
		fmt.Println("test check")
	} else if CheckOk(id, result) {
		fmt.Println("ok")
	} else {
		fmt.Println("fail")
	}
}
