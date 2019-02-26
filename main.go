package gojsonconf

import (
	"encoding/json"
	"os"
)

func GetConfig(path string, configuration interface{}) (err error) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	return
}
