package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GeneralRemoteCallTimeout = 10 // seconds
)

//=============================================================
//
//=============================================================

func RemoteCall(method string, url string, token string) (*http.Response, []byte, error) {
	request, err := http.NewRequest(method, url, nil)
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

func ParseJsonToMap(jsonByes []byte) (map[string]interface{}, error) {
	if jsonByes == nil {
		return nil, errors.New("jsonBytes can't be nil")
	}
	var v interface{}
	err := json.Unmarshal(jsonByes, &v)
	if err != nil {
		return nil, err
	}
	json_map, ok := v.(map[string]interface{})
	if !ok {
		return nil, errors.New("parse json error")
	}

	return json_map, nil
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

	var v interface{}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	m, ok := v.(map[string]interface{})
	if ok {
		return m, nil
	}

	return nil, fmt.Errorf("can't convert request.body to a map: %s.", string(data))
}
