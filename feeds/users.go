package feeds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sethgrid/pester"
)

// Example only configurable with env variables
//
// u := feeds.NewUser()
// go u.Start()
//
// The feed endpoint accept a parameter called "last" which is used as cursor for the current feed position
// This cursor will be saved when the update process is done
// GET http://example.de/feed?last=1503769504008
// { "items": [{ "name": "peter" }] }
//

var configuration *config

var (
	Feed            = "user"
	ErrUnMarshaling = errors.New("user feed: json parsing error")
	ErrClient       = errors.New("user feed: client connection error")
	ErrReadBody     = errors.New("user feed: response payload error")
)

type (
	config struct {
		URL              string `env:"FEED_USER_URL"`
		Table            string `env:"FEED_USER_TABLE" envDefault:"users"`
		Interval         int64  `env:"FEED_USER_INTERVAL"`
		ConnectionString int    `env:"FEED_USER_CONN_STRING"`
		Dialect          string `env:"FEED_USER_DIALECT" envDefault:"postgres"`
	}
	userFeed struct {
		Last   uint64
		ticker *time.Ticker
	}
	feedItem struct {
		Name string `json:"name"`
	}
	feed struct {
		Items []feedItem `json:"items"`
	}
)

// init loads the config from ENV
func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	err = godotenv.Load(path.Join(pwd, "feeds", ".env"))

	if err != nil {
		log.Print("No .env file could be found\n")
	}

	cfg := &config{}
	err = env.Parse(cfg)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	configuration = cfg
}

// Start the timer
func NewUser() *userFeed {
	u := &userFeed{}
	u.ticker = time.NewTicker(time.Duration(configuration.Interval) * time.Second)
	return u
}

// Poll starts the mechanism to initiate updates in certain intervals
func (u *userFeed) Start() {
	for _ = range u.ticker.C {
		u.poll(u.Last)
	}
}

func (u *userFeed) poll(last uint64) {
	feed, err := u.request()

	if err != nil {
		switch err.(type) {
		default:
			fmt.Println(err)
		}
	} else if err := u.update(feed); err == nil {
		u.saveLast(last)
	}

	fmt.Println(feed)
}

// update the database with the feed payload
func (u *userFeed) update(f *feed) error {
	return nil
}

// saveLast store the cursor of the current feed position
func (u *userFeed) saveLast(c uint64) {
	u.Last = c
}

// Request the feed endpoint and return the feed results
func (u *userFeed) request() (*feed, error) {
	client := pester.New()
	client.Concurrency = 3
	client.MaxRetries = 5
	client.Backoff = pester.ExponentialBackoff
	client.KeepLog = true

	resp, err := client.Get(configuration.URL)
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

	feed := &feed{}
	if err := json.Unmarshal(payload, feed); err != nil {
		return nil, ErrUnMarshaling
	}

	return feed, nil
}
