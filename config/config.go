package config

// Config 服务器配置
type Config struct {
	GRPCPort       string               `json:"GRPCPort"`       // grpc监听地址
	HTTPPort       string               `json:"HTTPPort"`       // http监听地址
}

func ReadConfig(file string, config interface{}) (err error) {
return nil
}
