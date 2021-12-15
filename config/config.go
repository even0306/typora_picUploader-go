package config

import (
	"encoding/json"
	"io/ioutil"
	"nextcloudUploader/logs"
	"os"
)

// type Configs map[string]json.RawMessage

// var configPath string = "./config.json"
var logging = logs.LogFile()

type picBed struct {
	Picbed string `json:"picBed"`
}

type nextcloud struct {
	picBed
	UploadUrl   string `json:"uploadUrl"`
	DownloadUrl string `json:"downloadUrl"`
	Path        string `json: "path"`
	User        string `json:"user"`
	Passwd      string `json:"passwd"`
	Proxy       string `json:"proxy"`
}

type aliyunOss struct {
	picBed
	Endpoint        string `json:"bucket"`
	BucketName      string `json:"bucketName"`
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Proxy           string `json:"proxy"`
}

// var Config struct {
// 	Bucket string
// 	Domain string
// 	User   string
// 	Passwd string
// 	Proxy  string
// }

var Config struct {
	PicBed     string
	Bucket     string
	Domain     string
	BucketName string
	Path       string
	User       string
	Passwd     string
	Proxy      string
}

func ReadConfig() interface{} {
	jsonFile, err := os.Open("E:/config.json")
	if err != nil {
		logging.Printf("打开配置文件失败，error：%v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logging.Printf("读取配置文件失败，error：%v", err)
	}

	var pb picBed
	json.Unmarshal([]byte(byteValue), &pb)
	getConfigValue(&byteValue, &pb)

	return Config
}

func getConfigValue(byteValue *[]byte, picbed *picBed) {
	switch {
	case picbed.Picbed == "nextcloud":
		var nextcloud nextcloud
		json.Unmarshal([]byte(*byteValue), &nextcloud)
		Config.PicBed = nextcloud.Picbed
		Config.Bucket = nextcloud.UploadUrl
		Config.Domain = nextcloud.DownloadUrl
		Config.Path = nextcloud.Path
		Config.User = nextcloud.User
		Config.Passwd = nextcloud.Passwd
		Config.Proxy = nextcloud.Proxy
	case picbed.Picbed == "aliyunOss":
		var aliyunOss aliyunOss
		json.Unmarshal([]byte(*byteValue), &aliyunOss)
		Config.PicBed = aliyunOss.Picbed
		Config.Bucket = aliyunOss.Endpoint
		Config.Domain = aliyunOss.Endpoint
		Config.BucketName = aliyunOss.BucketName
		Config.User = aliyunOss.AccessKeyId
		Config.Passwd = aliyunOss.AccessKeySecret
		Config.Proxy = aliyunOss.Proxy
	default:
		logging.Print("不支持的图床类型")
		os.Exit(-1)
	}
}

// func getConfigValue(byteValue *[]byte, picbed *picBed) {
// 	switch {
// 	case picbed.Picbed == "nextcloud":
// 		var nextcloud nextcloud
// 		json.Unmarshal([]byte(*byteValue), &nextcloud)
// 		Config = nextcloud
// 	case picbed.Picbed == "github":
// 		var github github
// 		json.Unmarshal([]byte(*byteValue), &github)
// 		Config = github
// 	}
// }

// var Conf *MainConfig
// var Confs Configs

// var instanceOnce sync.Once
// var logging = LogFile()

// //从配置文件中载入json字符串
// func LoadConfig(path string) (Configs, *MainConfig) {
// 	buf, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		logging.Panicln("load config conf failed: ", err)
// 	}
// 	mainConfig := &MainConfig{}
// 	err = json.Unmarshal(buf, mainConfig)
// 	if err != nil {
// 		logging.Panicln("decode config file failed:", string(buf), err)
// 	}
// 	allConfigs := make(Configs, 0)
// 	err = json.Unmarshal(buf, &allConfigs)
// 	if err != nil {
// 		logging.Panicln("decode config file failed:", string(buf), err)
// 	}

// 	return allConfigs, mainConfig
// }

// //初始化 可以运行多次
// func SetConfig(path string) {
// 	allConfigs, mainConfig := LoadConfig(path)
// 	configPath = path
// 	Conf = mainConfig
// 	Confs = allConfigs
// }

// // 初始化，只能运行一次
// func Init(path string) *MainConfig {
// 	if Conf != nil && path != configPath {
// 		logging.Printf("the config is already initialized, oldPath=%s, path=%s", configPath, path)
// 	}
// 	instanceOnce.Do(func() {
// 		allConfigs, mainConfig := LoadConfig(path)
// 		configPath = path
// 		Conf = mainConfig
// 		Confs = allConfigs
// 	})

// 	return Conf
// }

// //初始化配置文件 为 struct 格式
// func Instance() *MainConfig {
// 	if Conf == nil {
// 		Init(configPath)
// 	}
// 	return Conf
// }

// //初始化配置文件 为 map格式
// func AllConfig() Configs {
// 	if Conf == nil {
// 		Init(configPath)
// 	}
// 	return Confs
// }

// //获取配置文件路径
// func ConfigPath(path string) string {
// 	configPath = path
// 	return configPath
// }

// //根据key获取对应的值，如果值为struct，则继续反序列化
// func (cfg Configs) GetConfig(key string, config interface{}) error {
// 	c, ok := cfg[key]
// 	if ok {
// 		return json.Unmarshal(c, config)
// 	} else {
// 		logging.Fatalf("fail to get cfg with key: %s", key)
// 		return fmt.Errorf("fail to get cfg with key: %s", key)
// 	}
// }
