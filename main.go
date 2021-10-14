package main

import (
	"apphealthmonitor/files"
	"fmt"
)

var FILE_NAME = "config.yaml"

func main() {
	configs := files.LoadConfig(FILE_NAME)

	fmt.Println(configs)
}
