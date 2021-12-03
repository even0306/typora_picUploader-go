package uploadFile

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/op/go-logging"
)

//TimeoutDialer 连接超时和传输超时
func timeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

//上传接口，传url，文件二进制，参数头
func UploadFile(rURL *string, b *[]byte, header *map[string]string) error {
	var logging = logging.MustGetLogger("upload")
	req, err := http.NewRequest("PUT", *rURL, bytes.NewBuffer(*b))
	if err != nil {
		logging.Info(fmt.Sprintf("http newrequest error %s", err))
		return err
	}
	for h, v := range *header {
		req.Header.Set(h, v)
	}

	connectTimeout := 120 * time.Second
	readWriteTimeout := 5184000 * time.Millisecond
	client := http.Client{
		//忽略证书验证
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: false,
			},
			Dial: timeoutDialer(connectTimeout, readWriteTimeout),
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		logging.Info(fmt.Sprintf("http client error %s", err))
		return err
	}
	if resp != nil {
		// 判断请求状态
		if resp.StatusCode == http.StatusOK {
			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logging.Info(err)
				return err
			}
			logging.Info(fmt.Sprintf("\n【请求地址】： %s \n【请求参数】： %s \n【请求头】： %s \n【返回】 : %s \n",
				*rURL, "上传文件", *header, string(respData)))
			fmt.Println(string(respData))
			return nil
		} else if resp.StatusCode != http.StatusOK {
			respData, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				logging.Info(err)
				return err
			}

			logging.Info(fmt.Sprintf("\n【请求地址】： %s \n【请求参数】： %s \n【请求头】： %s \n【返回】 : %s \n",
				*rURL, "上传文件", *header, string(respData)))
			return errors.New("上传文件请求成功，上传成功")
		}
		return errors.New("请求失败")
	}
	defer resp.Body.Close()

	return errors.New("请求失败")
}
