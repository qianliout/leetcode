package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database    DBConfig    `yaml:"database"`
	CrawlConfig CrawlConfig `yaml:"crawl"`
}

type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

type CrawlConfig struct {
	Debug  bool     `yaml:"debug"`
	Period []string `yaml:"period"` // 爬取特定的周期
	Code   []string `yaml:"code"`   // 只爬取特定的
}

var conf *Config

func GetDefaultConfig() *Config {
	c := &Config{
		Database: DBConfig{
			Username: "root",
			Password: "root",
			Host:     "127.0.0.1",
			Port:     3306,
			DBName:   "stock",
		},
		CrawlConfig: CrawlConfig{
			Debug:  true,
			Period: []string{},
			Code:   []string{},
		},
	}
	return c
}

func SetDefaultConfig(c *Config) *Config {
	return c
}

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
