package feed

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/sethgrid/pester"
)

func NewPesterRequester() *PesterRequester {
	return &PesterRequester{}
}

type PesterRequester struct{}

func (r *PesterRequester) Request(url string) (*RootFeed, error) {
	client := pester.New()
	client.Concurrency = 3
	client.MaxRetries = 5
	client.Backoff = pester.ExponentialBackoff
	client.KeepLog = true

	resp, err := client.Get(url)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("error closing response body %q\n", err)
		}
	}()

	if err != nil {
		return nil, ErrClient
	}

	payload, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, ErrReadBody
	}

	f := &RootFeed{}
	if err := json.Unmarshal(payload, f); err != nil {
		return nil, ErrUnMarshaling
	}

	return f, nil

}
