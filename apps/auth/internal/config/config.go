// @Author huzejun 2024/3/24 21:52:00
package config

import (
	"flag"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name       string      `yaml:"name"`
	ServerID   int         `yaml:"server_id"`
	Port       int         `yaml:"port"`
	Log        string      `yaml:"log"`
	Etcd       *conf.Etcd  `yaml:"etcd"`
	Redis      *conf.Redis `yaml:"redis"`
	Mysql      *conf.Mysql `yaml:"mysql"`
	GrpcServer *conf.Grpc  `yaml:"grpc_server"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/auth.yaml", "config file")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
