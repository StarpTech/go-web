package user

import (
	"fmt"
	"time"

	"github.com/starptech/go-web/feed"
)

// Example:
// cfg := feed.Config{Table: "user", Interval: 5, URL: "http://example.de/feed"}
// store := feed.NewMemoryStore()
// pester := feed.NewPesterRequester()
// uf := userFeed.NewUser(cfg, store, pester)
// go uf.Start()
//
// The feed endpoint accept a parameter called "last" which is used as cursor for the current feed position
// This cursor will be saved when the update process is done
// GET http://example.de/feed?last=1503769504008
// { "items": [{ "name": "peter" }] }
//

type Feed struct {
	after     uint64
	ticker    *time.Ticker
	store     feed.FeedStore
	config    feed.Config
	requester feed.FeedRequester
}

// NewUser start the timer and return a new userFeed instance
func NewUser(config feed.Config, s feed.FeedStore, r feed.FeedRequester) *Feed {
	u := &Feed{}
	u.config = config
	u.store = s
	u.requester = r
	u.after = u.store.GetPosition(u.config.Table)
	u.ticker = time.NewTicker(time.Duration(u.config.Interval) * time.Second)
	return u
}

// Start poll updates in certain intervals
func (u *Feed) Start() {
	for range u.ticker.C {
		u.poll(u.after)
	}
}

// poll start a request against the feed endpoint
func (u *Feed) poll(after uint64) {
	feed, err := u.requester.Request(u.config.URL + "?after=" + string(u.after))

	if err != nil {
		switch err.(type) {
		default:
			fmt.Println(err)
		}
	} else if err := u.store.Save(feed); err == nil {
		if err := u.store.SetPosition(u.config.Table, after); err == nil {
			u.after = after
		}
	}

	fmt.Println(feed)
}
