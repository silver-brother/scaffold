package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func MustLoad(path string, result any) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error: config file %s, %s", path, err.Error())
	}
	err = yaml.Unmarshal(data, result)
	if err != nil {
		log.Fatalf("error: unmarshalling %s", err.Error())
	}
}
