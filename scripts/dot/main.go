package main

import (
	"os"

	"github.com/thanhpk/randstr"
)

func main() {
	token := ""

	for i := 0; i < 1; i++ {
		token = randstr.Hex(16)

	}

	dotEnv := os.WriteFile(".env", []byte(`NEXT_PUBLIC_BW_SECRET_KEY="`+token+`"`), 0644)

	if dotEnv != nil {
		panic(dotEnv)
	}
}
