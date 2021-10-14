package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	Sites    []string `yaml:sites,flow`
	Requests struct {
		Quantity     int `yaml:"quantity"`
		DelaySeconds int `yaml:"delaySeconds"`
	}
}

func LoadConfig(fileName string) Configs {
	configs := Configs{}

	filePath, _ := filepath.Abs(fileName)
	fileBuffer, err := ioutil.ReadFile(filePath)

	if err != nil {
		finishLoaderWithErr("Error in open configs file", err)
	}

	err = yaml.Unmarshal(fileBuffer, &configs)

	if err != nil {
		finishLoaderWithErr("Error in decode file", err)
	}

	return configs
}

func finishLoaderWithErr(msg string, err error) {
	fmt.Println(msg, err.Error())
	os.Exit(-1)
}
