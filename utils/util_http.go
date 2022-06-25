package utils

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// HTTPGet
func HTTPGet(url string) (res []byte, err error) {
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}

	req.SetRequestURI(url)
	err = fasthttp.DoTimeout(req, resp, 5*time.Second)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status_code != 200 (%d)", resp.StatusCode())
	}
	res = resp.Body()
	return
}
