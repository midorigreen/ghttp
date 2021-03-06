package ghttp

import (
	"fmt"
	"net/http"
)

func execGet(g GHttp) (body []byte, err error) {
	q := buildQuery(g.Query)
	resp, err := http.Get(buildGet(g, q))
	if err != nil {
		return body, err
	}
	defer resp.Body.Close()

	return branchStatusCode(*resp)
}

func buildQuery(m map[string]string) string {
	var q string
	for k, v := range m {
		q += fmt.Sprintf("%v=%v&", k, v)
	}
	return q[:len(q)-1]
}

func buildGet(g GHttp, q string) string {
	return fmt.Sprintf("%s?%s", g.getURL(), q)
}
