package global

var (
	// 并发请求chan
	CH_BULK = make(chan struct{}, 50)
)
