package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

var HttpClient = &http.Client{
	Timeout: 10 * time.Second,
}

/*
type Input struct {
	Url      string
	Method   string
	JsonData interface{}
}

*/
type Output struct {
	StatusCode int
	Body       interface{}
}

/*
func Do(input *Input, hasToken bool) (*http.Response, error) {
	var byte []byte
	var err error
	if input.JsonData != nil {
		byte, err = json.Marshal(input.JsonData)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(input.Method, input.Url, bytes.NewReader(byte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if hasToken {
		//application/x-www-form-urlencoded //application/json

		req.Header.Set(
			"Authorization",
			//"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODMyOTMzNjMsImlzcyI6Im11c2VuZXR3b3JrLm9yZyIsIlVJRCI6NX0.PokT-otwn4RlWl-eQ8S4ykwwxeCMBP00qfLVWu4uci0")
			//id=5
			//"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODQ1OTk3MDYsImlzcyI6Im11c2VuZXR3b3JrLm9yZyIsIlVJRCI6NTJ9.qpbw_MCtwLUTKC5DtGzAiZj5XpEM62GiO19CKK5of_4")
			//id=142
			"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU1NTcyNjEsImlzcyI6Im11c2VuZXR3b3JrLm9yZyIsIlVJRCI6MTQyfQ.PwYgQXokWrtGyjlhnOMItJhDHNr18GADNowc-svoDko")
	}
	//
	//req.Header.Set("Origin", "localhost")

	response, err := HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http请求错误：%s", err)
	}
	return response, nil
}

func HandleResponse(response *http.Response) (int, []byte, error) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return response.StatusCode, bodyBytes, err
}
*/

type Input struct {
	Path    string
	Method  string
	Payload interface{}
}

func Do(input *Input) ([]byte, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	var byte []byte
	var err error

	if input.Payload != nil {
		byte, err = json.Marshal(input.Payload)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(input.Method, input.Path, bytes.NewReader(byte))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := httpClient.Do(req)
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New(string(bodyBytes))
	}
	return bodyBytes, nil
}

func PostWithFormData(method, url string, postData *map[string]string) ([]byte, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return data, nil
}
