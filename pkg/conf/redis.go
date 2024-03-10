// @Author huzejun 2024/3/10 20:22:00
package conf

type Redis struct {
	Address  []string `yaml:"address"`
	Db       int      `yaml:"db"`
	Password string   `yaml:"password"`
	Prefix   string   `yaml:"prefix"`
}
