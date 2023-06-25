package main

import (
	"errors"
    "fmt"
	"log"
	"os"
	"path"
    "strings"
	"syscall"
)

func main() {
	p, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if !ok {
		p = path.Join(os.Getenv("HOME"), ".config")
	}
	p = path.Join(p, "browser-matcher/config.json")

    if len(os.Args) == 2 && strings.Contains(os.Args[1], "-h") {
        fmt.Printf("Usage: %s\n", os.Args[0])
        fmt.Println("Browser Matcher is a tool that automatically matches URLs to the appropriate web browser based on preconfigured patterns.")
        fmt.Println()
        fmt.Printf(
            "To use Browser Matcher, create a JSON configuration file at \"%s\" that specifies the web browsers you want to use and the rules for matching URLs to specific browsers.\n",
            p,
        )

        os.Exit(0)
    }

	errLog := log.New(os.Stderr, "", 0)

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
