package utils

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
)

var fileTypeMap sync.Map

func init() {
	fileTypeMap.Store("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif")  //GIF (gif)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc")  //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706d70", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")  //可执行文件
	fileTypeMap.Store("406563686f206f66660d", "bat")  //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")   //gz文件
	fileTypeMap.Store("504b0304140006000800", "docx") //docx文件
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov") //Quicktime (mov)
}

func checkType(fileCode *string) (fileType string) {
	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(*fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(*fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return
}

// 获取前面结果字节的二进制
func bytesToHexString(src *[]byte) string {
	res := bytes.Buffer{}
	if src == nil || len(*src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range *src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
}

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileExt(fSrc *[]byte) string {
	fileCode := bytesToHexString(fSrc)
	fileType := checkType(&fileCode)
	return fileType
}
