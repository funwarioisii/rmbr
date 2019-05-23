package main

import (
	"fmt"
	"rmbr/logic"
)

func main() {
	p := logic.StrStdin()
	p = logic.NewlineSeparate(p)
	p = logic.PeriodNewLine(p)
	fmt.Println(p)
}