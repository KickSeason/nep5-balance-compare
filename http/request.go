package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

//Request do a http request
func Request(method string, url string, reader io.Reader) ([]byte, error) {
	var request *http.Request
	if method == "POST" {
		var err error
		request, err = http.NewRequest("POST", url, reader)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		request.Header.Set("Content-Type", "application/json")
	} else {
		fmt.Println("not support http method")
		return nil, errors.New("not support")
	}

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return respBytes, nil
}
