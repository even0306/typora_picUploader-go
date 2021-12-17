package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"picUploader/logs"
	getexecpath "picUploader/utils/getExecPath"
)

var logging = logs.LogFile()

type picBed struct {
	Picbed string `json:"picBed"`
}

type nextcloud struct {
	picBed
	UploadUrl   string `json:"uploadUrl"`
	DownloadUrl string `json:"downloadUrl"`
	Path        string `json:"path"`
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
	jsonFile, err := os.Open(getexecpath.GetLocalPath() + "/config.json")
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
