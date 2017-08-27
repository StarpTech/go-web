package feed

import (
	"errors"
	"time"
)

var (
	ErrUnMarshaling = errors.New("feed: json parsing error")
	ErrClient       = errors.New("feed: client connection error")
	ErrReadBody     = errors.New("feed: response payload error")
)

type (
	Feed struct {
		after  uint64
		ticker *time.Ticker
	}
	RootFeed struct {
		Items []FeedItem `json:"items"`
	}
	Config struct {
		URL      string
		Table    string
		Interval int64
	}
	FeedItem      interface{}
	FeedRequester interface {
		Request(url string) (*RootFeed, error)
	}
)
