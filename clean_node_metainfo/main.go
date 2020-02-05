package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		println(fmt.Sprintf(`error retrieving working directory: %v`, err.Error()))
		return
	}
	cleanFolder(dir)
}
