package req

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Post请求
func Post(reqUrl string, proxy string, header map[string]string, data string) ([]byte, int, error) {
	client := &http.Client{}
	if proxy != "" {
		var proxyURI, _ = url.Parse(proxy)
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxyURI),
		}
	}
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(data))
	if err != nil {
		return nil, 0, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

//Get请求
func Get(reqUrl, proxy string, header map[string]string) ([]byte, int, error) {
	client := &http.Client{}
	if proxy != "" {
		var proxyURI, _ = url.Parse(proxy)
		client.Transport = &http.Transport{
			// 设置代理
			Proxy: http.ProxyURL(proxyURI),
		}
	}
	//提交请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, 0, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}

//Post请求
func PostForm(reqUrl string, postData map[string][]string) ([]byte, int, error) {
	client := &http.Client{}
	//post请求
	resp, err := client.PostForm(reqUrl, postData)
	if err != nil {
		return nil, 404, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}
	return body, resp.StatusCode, nil
}
