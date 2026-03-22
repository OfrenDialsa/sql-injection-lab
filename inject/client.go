package inject

import (
	"io"
	"net/http"
)

func send(baseURL, payload string) string {
	url := baseURL + "/user?id=" + payload

	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
