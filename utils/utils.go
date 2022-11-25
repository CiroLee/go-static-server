package utils

import (
	"fmt"
	"github.com/CiroLee/go-static-server/config"
	"os"
	"path"
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

func GetUrlByEnv(basePath, filename string) string {
	env, _ := GetEnv()
	if env.Mode == "debug" {
		return path.Join(fmt.Sprintf("%v:%v", config.DevHost, env.Port), basePath, filename)
	}
	return path.Join(config.ProdHost, basePath, filename)
}
