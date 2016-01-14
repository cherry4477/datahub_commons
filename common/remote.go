package common

import (
	"io/ioutil"
	"net/http"
	"time"
	"bytes"
	
	"github.com/asiainfoLDP/datahub_commons/log"
)

const (
	GeneralRemoteCallTimeout = 10 // seconds
)

//=============================================================
//
//=============================================================

func RemoteCallWithBody(method string, url string, token string, body []byte, contentType string) (*http.Response, []byte, error) {
	log.DefaultLogger().Debugf("method: %s, url: %s, token: %s, contentType: %s, body: %s", method, url, token, contentType, string(body))
	
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
	if contentType != "" {
		request.Header.Set("Content-Type", contentType)
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

func RemoteCallWithJsonBody(method string, url string, token string, jsonBody []byte) (*http.Response, []byte, error) {
	return RemoteCallWithBody(method, url, token, jsonBody, "application/json; charset=utf-8")
}

func RemoteCall(method string, url string, token string) (*http.Response, []byte, error) {
	return RemoteCallWithBody(method, url, token, nil, "")
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
	
	m, err := ParseJsonToMap(data)
	if err != nil {
		log.DefaultLogger().Debugf("ParseJsonToMap r.Body (%s) error: %s", string(data), err.Error())
	}
	
	return m, err
}
