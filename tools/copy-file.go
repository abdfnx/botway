package tools

import (
	"log"
	"os"
)

func Copy(src, dst string) {
	bytesRead, err := os.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(dst, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
