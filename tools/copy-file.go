package tools

import (
	"io/ioutil"
	"log"
)

func Copy(src, dst string) {
	bytesRead, err := ioutil.ReadFile(src)

	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(dst, bytesRead, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
