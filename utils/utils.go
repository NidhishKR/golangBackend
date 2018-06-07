package utils

import (
	"os"
	"strings"
)

func GetFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		panic("Please specify Environment")
	} else {
		filename := []string{"config.", env, ".json"}
		FilePath := strings.Join(filename, "")
		return FilePath
	}
}
