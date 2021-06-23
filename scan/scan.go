package scan

import (
	"fuzz/utils"
	"log"
	"strings"
	"time"
)

func Scan(urlAddress string, concurrency uint, timeOutUint uint, rules []string) {
	urlAddress = strings.Trim(urlAddress, "/") + "/"
	timeOut := time.Duration(timeOutUint) * time.Second

	// 创建http请求客户端
	httpClient := utils.NewHttpClient(timeOut, true)

	// 判断主站点是否正常访问
	_, err := httpClient.Get(urlAddress)
	if err != nil {
		log.Fatalf("该站点无法访问：%s", err)
	}




}
