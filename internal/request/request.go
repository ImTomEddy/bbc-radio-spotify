package request

import (
	"io/ioutil"
	"net/http"
)

// MakeGetRequest makes a HTTP Get request and returns the content as a byte slice
func MakeGetRequest(requestURI string) ([]byte, error) {
	response, err := http.Get(requestURI)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, err
}
