package main

import (
	"fmt"

	"github.com/thanhpk/randstr"
)

func main() {
	token := ""

	for i := 0; i < 1; i++ {
		token = randstr.Hex(16)

	}

	fmt.Println(`NEXT_PUBLIC_BW_SECRET_KEY="`+token+`"`)
}
