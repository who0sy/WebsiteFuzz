package utils

func GetRandomUa() string {

	// 该包需要连接github cache, 因网络原因故使用固定UA
	//return userAgent.Random()
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"

}
