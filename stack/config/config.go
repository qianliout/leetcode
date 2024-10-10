package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database     DBConfig     `yaml:"database"`
	SpiderConfig SpiderConfig `yaml:"spiderConfig"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

type SpiderConfig struct {
	Debug bool `yaml:"debug"`
}

var conf *Config

func ParseConfig(str string) (*Config, error) {
	if conf != nil {
		return conf, nil
	}
	// 读取YAML文件
	yamlFile, err := os.ReadFile(str)
	if err != nil {
		return nil, err
	}

	// 解析YAML内容到Config结构体
	var c Config
	if err := yaml.Unmarshal(yamlFile, &c); err != nil {
		return nil, err
	}
	conf = &c
	return conf, nil
}
