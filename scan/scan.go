package scan

import (
	"fmt"
	"fuzz/utils"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	NotFoundContent    string
	NotFoundStatusCode int
)

func Scan(urlAddress string, concurrency int, timeOutUint int, rules []string) {
	urlAddress = strings.Trim(urlAddress, "/") + "/"

	// 创建http请求客户端
	httpClient := utils.NewHttpClient(time.Duration(timeOutUint)*time.Second, true)

	// 判断主站点是否正常访问
	statusCode, _, _ := httpClient.Request("GET", urlAddress, nil)
	if statusCode == 0 {
		log.Fatalf("该站点无法访问！")
	}

	// 构造随机字符串路径获取404页面地址，清理误报
	randUrl := fmt.Sprintf("%s/%s", urlAddress, utils.RandString(10))
	statusCode, _, body := httpClient.Request("GET", randUrl, nil)
	NotFoundStatusCode = statusCode
	NotFoundContent = body

	threadChan := make(chan struct{}, concurrency)
	defer close(threadChan)
	wg := &sync.WaitGroup{}
	begin := time.Now()
	for i := 0; i < len(rules); i++ {
		threadChan <- struct{}{}
		url := urlAddress + rules[i]
		wg.Add(1)
		go func() {
			defer wg.Done()
			statusCode, _, content := httpClient.Request("GET", url, nil)
			if statusCode == 0 {
				return
			}
			// 暂定状态码为（200， 400）之间或状态码为403的为正常页面
			if (statusCode >= 200 && statusCode < 400) || statusCode == 403 {
				// 计算和404页面相似性
				isSimilar, _ := utils.GetSimilar(NotFoundContent, content)
				if !isSimilar || statusCode != NotFoundStatusCode {
					log.Printf("【%d】: %s", statusCode, url)
				}
			}
			<-threadChan
		}()
	}
	wg.Wait()
	fmt.Printf("扫描完成，扫描时间为: %fs\n", time.Now().Sub(begin).Seconds())
}
