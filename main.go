package main

import (
	"fmt"
	"os"

	"github.com/hedwig100/lettera/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
