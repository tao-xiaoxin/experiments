package pkg1

import (
	"fmt"
	_ "gitee.com/tao-xiaoxin/study-basic-go/syntax/prog-init-order/pkg3"
)

var (
	_ = constInitCheck()
	v = variableInit("v")
)

const (
	c = "c"
)

func constInitCheck() string {
	if c != "" {
		fmt.Println("pkg1: const c has been initialized!")
	}

	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg1: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg1: init function invoked")
}
