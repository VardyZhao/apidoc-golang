package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Conf map[string]interface{} `yaml:"config"`
}

func Load(fileName string) *Config {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var c Config
	c.Conf = make(map[string]interface{})
	err = yaml.Unmarshal(data, c.Conf)
	if err != nil {
		panic(err)
	}

	return &c
}

func (c *Config) Get(key string) (interface{}, error) {
	keys := strings.Split(key, ".")
	v := c.Conf

	for i := 0; i < len(keys); i++ {
		val, ok := v[keys[i]]
		if !ok {
			return nil, errors.New("key not found")
		}

		// 如果是最后一个键，直接返回值
		if len(keys)-1 == i {
			return val, nil
		}

		// 如果不是最后一个键，检查是否为 map[string]interface{} 类型
		v, ok = val.(map[string]interface{})
		if !ok {
			return nil, errors.New("value is not a map")
		}
	}

	return nil, nil
}

func (c *Config) GetString(key string) (string, error) {
	val, err := c.Get(key)
	if err != nil {
		return "", errors.New("Error happen for " + key + ", " + err.Error())
	}
	var res string
	switch v := val.(type) {
	case int:
		res = strconv.Itoa(v) // 将整数转换为字符串
	case string:
		res = v
	default:
		return "", errors.New("Unsupported type for " + key)
	}
	return res, nil
}

func (c *Config) GetInt(key string) (int, error) {
	val, err := c.Get(key)
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}
