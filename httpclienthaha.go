package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	_ "net/http/pprof"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

var jobDone = make(chan bool, 1)

func main() {
	size, err := GetMaxFileDescriptorSize()
	queueCh := make(chan bool, size)
	wg := sync.WaitGroup{}
	file, err := os.Open("./ports.log")

	if err != nil {
		panic(fmt.Sprintf("open logfile: %s\n", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	defaultTransport = &http.Transport{
		Proxy: proxy,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: -1,
		}).DialContext,
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: -1,
		MaxConnsPerHost:     0,
		IdleConnTimeout:     5 * time.Second,
		DisableCompression:  true,
	}
	defaultClient = &http.Client{Transport: defaultTransport, Timeout: time.Duration(5) * time.Second}

	go http.ListenAndServe(":8080", nil)

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}

		wg.Add(1)
		ip, port := ParseURLToIPAndPort(scanner.Text())
		queueCh <- true
		go func(ip string, port string) {
			defer wg.Done()
			defer func() {
				<-queueCh
			}()
			body, err := RequestBody(ip, port)
			if err != nil {
				fmt.Printf("%s:%s err:%s\n", ip, port, err)
				return
			}

			if strings.Contains(body, "EasouSpider") {
				fmt.Printf("%s:%s passed\n", ip, port)
			} else {
				fmt.Printf("%s:%s failed\n", ip, port)
			}
		}(ip, port)

	}
	wg.Wait()
	fmt.Println("done")

}

type contextKey int

const (
	proxyURLKey contextKey = 0
)

func proxy(req *http.Request) (*url.URL, error) {
	iproxy := req.Context().Value(proxyURLKey)
	if iproxy == nil {
		// fmt.Printf("no proxy found\n")
		return nil, nil
	}
	proxyURL := iproxy.(*url.URL)
	// fmt.Printf("proxy found:%v\n", proxyURL)
	return proxyURL, nil
}

var defaultTransport *http.Transport
var defaultClient *http.Client

func RequestBody(ip string, port string) (body string, err error) {

	proxyURL, _ := url.Parse(fmt.Sprintf("http://%s:%s", ip, port))
	ctx := context.WithValue(context.Background(), proxyURLKey, proxyURL)
	req, err := http.NewRequestWithContext(ctx, "GET", "http://baidu.com/robots.txt", nil)

	// req, err := http.NewRequest("GET", "http://baidu.com/robots.txt", nil)
	req.Header.Set("Connection", "close")
	req.Close = true
	resp, err := defaultClient.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	byte, err := ioutil.ReadAll(resp.Body)
	body = string(byte)
	return
}

func ParseURLToIPAndPort(url string) (string, string) {
	url = strings.Replace(url, ",", ":", -1)
	splits := strings.SplitN(url, ":", 2)
	return splits[0], splits[1]
}

// GetMaxFileDescriptorSize 获取操作系统最大打开文件数
func GetMaxFileDescriptorSize() (size int, err error) {
	out, err := exec.Command("/bin/bash", "-c", "ulimit -n").Output()
	if err != nil {
		return
	}
	size, err = strconv.Atoi(strings.Replace(string(out), "\n", "", -1))
	if err != nil {
		return
	}
	size = size/2 - 20
	return
}
