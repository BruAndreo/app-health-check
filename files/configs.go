package files

import (
	"fmt"
	"io/ioutil"
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

func LoadConfig() {
	configs := Configs{}

	filePath, _ := getFilePath()
	fileBuffer, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error in load config.yaml", err.Error())
	}

	err2 := yaml.Unmarshal(fileBuffer, &configs)

	if err2 != nil {
		fmt.Println("Error in decode file", err.Error())
	}

	fmt.Println("Seráá", configs)
}

func getFilePath() (string, error) {
	return filepath.Abs("config.yaml")
}
