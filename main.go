package main

import (
	"apphealthmonitor/files"
)

func main() {
	getConfigs()
}

func getConfigs() {
	files.LoadConfig()
}
