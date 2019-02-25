package gojsonconf

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJsonConfig(path string, t interface{}) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
