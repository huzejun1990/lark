// @Author huzejun 2024/3/7 1:32:00
package config

import (
	"flag"
	"lark/pkg/conf"
	"lark/pkg/utils"
)

type Config struct {
	Name     string      `yaml:"name"`
	ServerID int         `yaml:"server_id"`
	Port     int         `yaml:"port"`
	Log      string      `yaml:"log"`
	Etcd     *conf.Etcd  `yaml:"etcd"`
	Redis    *conf.Redis `yaml:"redis"`
}

var (
	config = new(Config)
)

var (
	confFile = flag.String("cfg", "./configs/api_gateway.yaml", "config file")
	serverId = flag.Int("sid", 1, "server id")
	port     = flag.Int("p", 8088, "api gateway default listen port 8088")
)

func init() {
	flag.Parse()
	utils.YamlToStruct(*confFile, config)

	config.ServerID = *serverId
	config.Port = *port

}

func NewConfig() *Config {
	return config
}

func GetConfig() *Config {
	return config
}
