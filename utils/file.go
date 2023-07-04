package util

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(i string, name string) {
	f, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(i)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
