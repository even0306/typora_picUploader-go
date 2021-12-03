package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

// 判断数据类型
func FileType(file *string) (filetype string) {
	if *file != "" {
		req := strings.Split(*file, ":")[0]
		i := []byte(req)
		if req == "data" {
			filetype = "base64"
		} else if req == "http" || req == "https" {
			filetype = "url"
		} else if (i[0] >= 65 && i[0] <= 90) || (i[0] >= 97 && i[0] <= 122) {
			filetype = "local"
		} else {
			fmt.Printf("无法识别的类型")
		}

	} else {
		fmt.Printf("数据不能为空")
	}
	return
}

//读取文件为二进制格式
func ReadFile(path *string) (b []byte, e error) {
	file, err := os.Open(*path)
	if err != nil {
		fmt.Printf("open file failed, error: %v", err)
		return
	}
	defer file.Close()
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("读取文件失败，error: %v", err)
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return chunks, err
}

// 下载文件
func DownloadFile(imgUrl *string, path *string, proxy *string) {
	timeout := time.Duration(30 * time.Second)
	pxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(*proxy)
	}
	transport := &http.Transport{Proxy: pxy}
	client := http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
	resp, err := client.Get(*imgUrl)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()

	// 创建一个文件用于保存
	out, err := os.Create(*path)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

// 创建UUID作为文件名
func CreateUUID() (key string) {
	uuid := uuid.New()
	key = uuid.String()
	return
}

// 获取执行文件当前所在路径
func GetLocalPath() (exPath string) {
	ex, err := os.Executable()
	if err != nil {
		fmt.Printf("获取路径失败，error：%v", err)
	}
	exPath = filepath.Dir(ex)
	return
}
