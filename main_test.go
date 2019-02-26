package gojsonconf

import "testing"

type Config struct {
	Port string
}

func TestGetConfig(t *testing.T) {
	got := Config{}
	want := Config{Port: "3000"}
	GetConfig("./conf.json", &got)

	if got != want {
		t.Errorf("want '%v' but got '%v'", want, got)
	}
}
