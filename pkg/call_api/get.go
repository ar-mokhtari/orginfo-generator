package call_api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func CallAPIs(httpMethod string, URL string, Params map[string]string, Headers map[string]string, Body []byte) (res []byte, err error) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(httpMethod, URL, bytes.NewBuffer(Body))
	if err != nil {
		return nil, err
	}
	urlParams := url.Values{}
	for key, value := range Params {
		urlParams.Add(key, value)
	}
	for key, value := range Headers {
		req.Header.Set(key, value)
	}
	req.URL.RawQuery = urlParams.Encode()
	response, callErr := client.Do(req)
	if callErr != nil {
		fmt.Println(callErr.Error())
		os.Exit(1)
		return nil, callErr
	}
	defer response.Body.Close()
	res, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		log.Fatalln(readErr)
		return nil, readErr
	}
	return res, nil
}
