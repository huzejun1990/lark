// @Author huzejun 2024/3/10 20:27:00
package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func YamlToStruct(file string, out interface{}) (err error) {
	var buf []byte
	buf, err = ioutil.ReadFile(file)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, out)
	return
}
