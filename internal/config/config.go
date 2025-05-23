package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func ReadFile(cfg interface{}) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}
