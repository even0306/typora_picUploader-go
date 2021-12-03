package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

type Configs map[string]json.RawMessage

var configPath string = "./config.json"

type MainConfig struct {
	UploadUrl   string `json:"uploadUrl"`
	DownloadUrl string `json:"downloadUrl"`
	User        string `json:"user"`
	Passwd      string `json:"passwd"`
	Proxy       string `json:"proxy"`
}

var Conf *MainConfig
var Confs Configs

var instanceOnce sync.Once

//从配置文件中载入json字符串
func LoadConfig(path string) (Configs, *MainConfig) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}
	mainConfig := &MainConfig{}
	err = json.Unmarshal(buf, mainConfig)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	allConfigs := make(Configs, 0)
	err = json.Unmarshal(buf, &allConfigs)
	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}

	return allConfigs, mainConfig
}

//初始化 可以运行多次
func SetConfig(path string) {
	allConfigs, mainConfig := LoadConfig(path)
	configPath = path
	Conf = mainConfig
	Confs = allConfigs
}

// 初始化，只能运行一次
func Init(path string) *MainConfig {
	if Conf != nil && path != configPath {
		log.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
	}
	instanceOnce.Do(func() {
		allConfigs, mainConfig := LoadConfig(path)
		configPath = path
		Conf = mainConfig
		Confs = allConfigs
	})

	return Conf
}

//初始化配置文件 为 struct 格式
func Instance() *MainConfig {
	if Conf == nil {
		Init(configPath)
	}
	return Conf
}

//初始化配置文件 为 map格式
func AllConfig() Configs {
	if Conf == nil {
		Init(configPath)
	}
	return Confs
}

//获取配置文件路径
func ConfigPath(path string) string {
	configPath = path
	return configPath
}

//根据key获取对应的值，如果值为struct，则继续反序列化
func (cfg Configs) GetConfig(key string, config interface{}) error {
	c, ok := cfg[key]
	if ok {
		return json.Unmarshal(c, config)
	} else {
		return fmt.Errorf("fail to get cfg with key: %s", key)
	}
}
