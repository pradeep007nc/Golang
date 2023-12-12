package utils

import (
	"io/ioutil"
	"net/http"
)

func GetIpAddress() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
