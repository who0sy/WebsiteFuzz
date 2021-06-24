package main

import (
	"flag"
	"fuzz/scan"
	"fuzz/utils"
	"log"
	"net/url"
	"runtime"
)

var (
	urlAddress  string
	dictPath    string
	concurrency int
	timeOut     int
)

func init() {
	flag.StringVar(&urlAddress, "u", "", "目标站点地址")
	flag.StringVar(&dictPath, "d", "./dict/dict.txt", "爆破字典路径")
	flag.IntVar(&concurrency, "c", runtime.GOMAXPROCS(runtime.NumCPU()), "并发数")
	flag.IntVar(&timeOut, "t", 3, "单个请求规则超时时间")
}

func main() {
	flag.Parse()

	// 解析website
	_, err := url.ParseRequestURI(urlAddress)
	if err != nil {
		log.Fatalf("【%s】不是一个合法的URL", urlAddress)
	}

	// 解析字典
	rules, err := utils.LoadDictFile(dictPath)
	if err != nil {
		log.Fatalf("爆破字典【%s】加载失败：%s", dictPath, err)
	}

	log.Printf("扫描目标站点：【%s】", urlAddress)
	log.Printf("加载爆破文件路径：【%s】", dictPath)
	log.Printf("爆破规则数：【%d】", len(rules))
	log.Printf("并发数：【%d】", concurrency)
	log.Printf("请求超时时间：【%d】", timeOut)

	log.Println("开始扫描.......")
	log.Println("----------------------------------------------")

	// 开始扫描
	scan.Scan(urlAddress, concurrency, timeOut, rules)

}
