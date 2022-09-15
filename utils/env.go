package utils

import (
	"os"
	"path/filepath"
	"strings"
)

var devEnv = false

func init() {
	devEnv = os.Getenv("ge_dev_env") == "1"
}

func IsDevEnv() bool {
	return devEnv
}

func BinaryPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}
