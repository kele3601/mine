package app

import (
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
)

type Config struct {
	APP struct {
		Mode string `yaml:"mode"`
	}
	DB struct {
		Type string `yaml:"type"`
		DSN  string `yaml:"dsn"`
	}
	WEB struct {
		Port string `yaml:"port"`
	}
}

var config *Config

func NewConfig(path string) (*Config, error) {
	if nil != config {
		return config, nil
	}
	// 获取配置文件目录
	dir := filepath.Dir(path)
	viper.AddConfigPath(dir)
	// 获取配置文件名
	base := filepath.Base(path)
	viper.SetConfigName(base[:strings.LastIndex(base, ".")])
	// 获取配置文件类型
	ext := filepath.Ext(path)
	viper.SetConfigType(ext[1:])

	// 读取配置文件
	if err := viper.ReadInConfig(); nil != err {
		return nil, err
	}
	// 解析配置文件
	if err := viper.Unmarshal(&config); nil != err {
		return nil, err
	}
	return config, nil
}
