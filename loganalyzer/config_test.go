package loganalyzer

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	config := loadConfig()
	fmt.Println(config.LogPattern)
	for _, f := range config.LogFormat {
		fmt.Println(f)
	}
	fmt.Println(config.TimeFormat)
}
