# WebsiteFuzz

Golang版对目标站点进行目录和文件爆破, 实现相似性汉明距离从而对404相似页面误报清理。

## 使用说明

```shell
go build
./fuzz -u http://baidu.com -d ./dict/dict.txt -t 3 -c 20

# 查看帮助
./fuzz -h
```

![scan.png](./scan.png)

