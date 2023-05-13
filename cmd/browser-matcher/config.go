package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"regexp"
)

type BrowserId string

type Browser struct {
	Id   BrowserId `json:"id"`
	Bin  string    `json:"bin"`
	Args []string  `json:"args,omitempty"`
}

type Regexp struct {
	*regexp.Regexp
}

type Rule struct {
	Value   *Regexp   `json:"value"`
	Browser BrowserId `json:"browser"`
}

type Config struct {
	Browsers       []Browser `json:"browsers"`
	Rules          []Rule    `json:"rules"`
	DefaultBrowser BrowserId `json:"defaultBrowser"`
}

var ErrNoMatchFound = errors.New("config: no match found")
var ErrBrowserNotFound = errors.New("config: browser not found")

func (c Config) Match(url string) (BrowserId, error) {
	for _, r := range c.Rules {
		if r.Value.Match([]byte(url)) {
			return r.Browser, nil
		}
	}

	return "", ErrNoMatchFound
}

func (c Config) BrowserById(bId BrowserId) (Browser, error) {
	for _, b := range c.Browsers {
		if b.Id == bId {
			return b, nil
		}
	}

	return Browser{}, ErrBrowserNotFound
}

func (r *Regexp) UnmarshalText(b []byte) error {
	v, err := regexp.Compile(string(b))
	if err != nil {
		return err
	}

	r.Regexp = v

	return nil
}

func (r *Regexp) MarshalText() ([]byte, error) {
	if r.Regexp != nil {
		return []byte(r.Regexp.String()), nil
	}

	return nil, nil
}

func readConfig(p string) (Config, error) {
	f, err := os.Open(p)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}
