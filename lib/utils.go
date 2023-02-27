package lib

import (
	"fmt"
)

func P(a ...interface{}) {
	fmt.Println(a...)
}

func Version() string {
	return "0.1"
}