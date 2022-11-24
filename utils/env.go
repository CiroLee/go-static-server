package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	Mode string
	Port string
	Desc string
}

func GetEnv() (Env, error) {
	env := os.Getenv("GODAILYLIB_ENV")
	if env == "" {
		env = "development"
	}
	// 根据环境变量读取配置
	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatal(err)
	}
	// 读取默认配置
	err = godotenv.Load()
	if err != nil {
		return Env{}, nil
	}

	fmt.Printf("evn desc: %v\n", os.Getenv("desc"))

	return Env{
		Mode: os.Getenv("mode"),
		Port: os.Getenv("port"),
		Desc: os.Getenv("desc"),
	}, nil

}
