package files

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Configs struct {
	// sites    []string
	requests struct {
		quantity     int `yaml:"quantity"`
		delaySeconds int `yaml:"delaySeconds"`
	}
}

func LoadConfig() {
	configs := Configs{}

	filePath, _ := filepath.Abs("config.yaml")
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
