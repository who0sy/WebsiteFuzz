# WebsiteFuzz

Golang版对目标站点进行目录和文件爆破

## 使用说明

```shell
go build
./fuzz -u http://baidu.com -d ./dict/dict.txt -t 3

# 查看帮助
./fuzz -h
```