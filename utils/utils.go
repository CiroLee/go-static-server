package utils

import (
	"fmt"
	"github.com/CiroLee/go-static-server/config"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetUrlByEnv(filePath, filename string) string {
	env, _ := GetEnv()
	if env.Mode == "debug" {
		return fmt.Sprintf("%v:%v%v/%v", config.DevHost, env.Port, filePath, filename)
	}
	return fmt.Sprintf("%v%v/%v", config.ProdHost, filePath, filename)
}
