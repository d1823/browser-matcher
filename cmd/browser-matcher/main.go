package main

import (
	"errors"
	"log"
	"os"
	"path"
	"syscall"
)

func main() {
	errLog := log.New(os.Stderr, "", 0)

	p, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if !ok {
		p = path.Join(os.Getenv("HOME"), ".config")
	}
	p = path.Join(p, "browser-matcher/config.json")

	c, err := readConfig(p)
	if err != nil {
		errLog.Fatalln(err)
	}

	var url string
	if len(os.Args) == 2 {
		url = os.Args[1]
	}

	bId, err := c.Match(url)
	if errors.Is(err, ErrNoMatchFound) {
		bId = c.DefaultBrowser
	} else if err != nil {
		errLog.Fatalln(err)
	}

	b, err := c.BrowserById(bId)
	if err != nil {
		errLog.Fatalln(err)
	}

	args := []string{path.Base(b.Bin)}
	args = append(args, b.Args...)
	if len(url) > 0 {
		args = append(args, url)
	}

	err = syscall.Exec(b.Bin, args, os.Environ())
	if err != nil {
		errLog.Fatalln(err)
	}
}
