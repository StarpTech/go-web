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
	FeedItem interface{}
	RootFeed struct {
		Items []FeedItem `json:"items"`
	}
	Config struct {
		URL      string
		Table    string
		Interval int64
	}
)
