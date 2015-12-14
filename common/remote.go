package common

import (
	"io/ioutil"
	"net/http"
	"time"
	"bytes"
)

const (
	GeneralRemoteCallTimeout = 10 // seconds
)

//=============================================================
//
//=============================================================

func RemoteCallWithBody(method string, url string, token string, body []byte) (*http.Response, []byte, error) {
	var request *http.Request
	var err error
	if len(body) == 0 {
		request, err = http.NewRequest(method, url, nil)
	} else {
		request, err = http.NewRequest(method, url, bytes.NewReader(body))
	}
	if err != nil {
		return nil, nil, err
	}
	if token != "" {
		request.Header.Set("Authorization", token)
	}
	client := &http.Client{
		Timeout: time.Duration(GeneralRemoteCallTimeout) * time.Second,
	}
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}
	if err != nil {
		return nil, nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	return response, bytes, err
}

func RemoteCall(method string, url string, token string) (*http.Response, []byte, error) {
	return RemoteCallWithBody(method, url, token, nil)
}

func GetRequestData(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}

	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseRequestJsonAsMap(r *http.Request) (map[string]interface{}, error) {
	data, err := GetRequestData(r)
	if err != nil {
		return nil, err
	}
	
	return ParseJsonToMap(data)
}
