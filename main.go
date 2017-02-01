package deptest

import (
	"fmt"
	s "github.com/naoina/go-stringutil"
)

func Hello() {
	fmt.Println(s.ToSnakeCase("HelloWorld"))
}
