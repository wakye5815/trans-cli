package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/sling"
)

type doer struct{}

func (this *doer) Do(request *http.Request) (*http.Response, error) {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || 299 < response.StatusCode {
		defer response.Body.Close()
		byteArray, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Receive failed http response (status=%s, body=%v)", response.Status, string(byteArray))
	}

	return response, nil
}

func NewCustomSling() *sling.Sling {
	return sling.New().Doer(&doer{})
}
