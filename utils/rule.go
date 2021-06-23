package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func LoadDictFile(dictPath string) (rules []string, err error) {
	file, err := os.Open(dictPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	newReader := bufio.NewReader(file)
	for {
		rule, err := newReader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("爆破字典解析失败: 【%s】", err)
		}
		// 读取结束
		if err == io.EOF {
			break
		}
		rules = append(rules, strings.Trim(rule, "\r\n"))
	}

	if len(rules) == 0 {
		log.Fatalf("未读取到该文件【%s】的爆破规则，请重新输入！", dictPath)
	}

	return rules, nil

}
