package run

import (
	"encoding/base64"
	"fmt"
	"nextcloudUploader/uploadFile"
	"nextcloudUploader/utils"
	"os"
	"strings"
)

type Base64 struct {
	Data
}

type Local struct {
	Data
}

type Http struct {
	Data
}

type Data struct {
	filePath    string
	fileName    string
	UploadUrl   string
	DownloadUrl string
	Auth        map[string]string
	Proxy       string
	ConfigPath  string
}

var resq struct {
	fmtUrl string
	upn    string
}

type Upload interface {
	upload(uploadData string) string
}

func (b *Base64) upload(args string) string {
	utils.Init(b.ConfigPath)
	b.UploadUrl = utils.Conf.UploadUrl
	b.DownloadUrl = utils.Conf.DownloadUrl
	user := utils.Conf.User
	passwd := utils.Conf.Passwd
	b.Auth = map[string]string{"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+passwd))}
	b.filePath = strings.Split(strings.Split(args, "base64,")[1], ")")[0]
	file, err := base64.StdEncoding.DecodeString(string(b.filePath))
	if err != nil {
		fmt.Printf("解密base64失败，error: %v", err)
	}

	//判断文件格式
	filetype := utils.GetFileType(&file)
	if filetype == "" {
		fmt.Printf("没有匹配到该文件格式")
		os.Exit(3)
	}

	b.fileName = utils.CreateUUID() + "." + filetype
	resq.upn = b.UploadUrl + b.fileName
	err = uploadFile.UploadFile(&resq.upn, &file, &b.Auth)
	if err != nil {
		resq.fmtUrl = b.DownloadUrl + b.fileName + "\n"
	}
	return resq.fmtUrl
}

func (l *Local) upload(args string) string {
	utils.Init(l.ConfigPath)
	l.UploadUrl = utils.Conf.UploadUrl
	l.DownloadUrl = utils.Conf.DownloadUrl
	user := utils.Conf.User
	passwd := utils.Conf.Passwd
	l.Auth = map[string]string{"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+passwd))}
	l.filePath = args
	file, err := utils.ReadFile(&l.filePath)
	if err != nil {
		fmt.Printf("读取文件失败，error：%v", err)
	}

	//判断文件格式
	filetype := utils.GetFileType(&file)
	if filetype == "" {
		fmt.Printf("没有匹配到该文件格式")
		os.Exit(3)
	}

	l.fileName = utils.CreateUUID() + "." + filetype
	resq.upn = l.UploadUrl + l.fileName
	err = uploadFile.UploadFile(&resq.upn, &file, &l.Auth)
	if err != nil {
		resq.fmtUrl = l.DownloadUrl + l.fileName + "\n"
	}
	return resq.fmtUrl
}

func (h *Http) upload(args string) string {
	utils.Init(h.ConfigPath)
	h.UploadUrl = utils.Conf.UploadUrl
	h.DownloadUrl = utils.Conf.DownloadUrl
	user := utils.Conf.User
	passwd := utils.Conf.Passwd
	h.Auth = map[string]string{"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(user+":"+passwd))}
	h.Proxy = utils.Conf.Proxy
	tmp := utils.GetLocalPath() + "/tmp"
	h.filePath = args
	utils.DownloadFile(&h.filePath, &tmp, &h.Proxy)
	file, err := utils.ReadFile(&tmp)
	if err != nil {
		fmt.Printf("读取文件失败，error：%v", err)
	}

	//判断文件格式
	filetype := utils.GetFileType(&file)
	if filetype == "" {
		fmt.Printf("没有匹配到该文件格式")
		os.Exit(3)
	}

	h.fileName = utils.CreateUUID() + "." + filetype
	err = os.Remove(tmp)
	if err != nil {
		fmt.Printf("删除缓存图片失败，error：%v", err)
	}
	resq.upn = h.UploadUrl + h.fileName
	err = uploadFile.UploadFile(&resq.upn, &file, &h.Auth)
	if err != nil {
		resq.fmtUrl = h.DownloadUrl + h.fileName + "\n"
	}
	return resq.fmtUrl
}

func Run(up Upload, args *string) string {
	arg := up.upload(*args)
	return arg
}
