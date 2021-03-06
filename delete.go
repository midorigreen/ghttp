package ghttp

import (
	"net/http"
	"time"
)

func execDelete(g GHttp) (body []byte, err error) {
	tr := &http.Transport{
		IdleConnTimeout: 30 * time.Second,
	}
	client := http.Client{
		Transport: tr,
	}
	resp, err := client.Do(&http.Request{
		Method:     DELETE,
		RequestURI: g.getURL(),
	})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return branchStatusCode(*resp)
}
