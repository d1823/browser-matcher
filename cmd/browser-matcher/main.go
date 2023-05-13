package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"syscall"
)

func main() {
	p, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if !ok {
		p = path.Join(os.Getenv("HOME"), ".config")
	}
	p = path.Join(p, "browser-matcher/config.json")

	c, err := readConfig(p)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var url string
	if len(os.Args) == 2 {
		url = os.Args[1]
	}

	bId, err := c.Match(url)
	if errors.Is(err, ErrNoMatchFound) {
		bId = c.DefaultBrowser
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	b, err := c.BrowserById(bId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args := []string{path.Base(b.Bin)}
	args = append(args, b.Args...)
	if len(url) > 0 {
		args = append(args, url)
	}

	err = syscall.Exec(b.Bin, args, os.Environ())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
