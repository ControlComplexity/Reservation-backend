package config

import (
	"github.com/spf13/viper"
	"bytes"
	"io/ioutil"
	"os"
)

// Config 服务器配置
type Config struct {
	GRPCPort       string               `json:"GRPCPort"`       // grpc监听地址
	HTTPPort       string               `json:"HTTPPort"`       // http监听地址
}

func ReadConfig(file string, config interface{}) (err error) {
	viper.SetConfigType("yaml")
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	// 替换环境变量
	configRaw := []byte(os.ExpandEnv(string(content)))
	if err = viper.ReadConfig(bytes.NewBuffer(configRaw)); err != nil {
		return
	}
	return viper.Unmarshal(config)
}
