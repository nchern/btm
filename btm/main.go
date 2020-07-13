package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	srcURL = "https://storage.googleapis.com/samofly/track/index.html"

	waitTimeout = 1 * time.Minute
)

func fetch() (io.ReadCloser, error) {
	resp, err := http.Get(srcURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 200 {
		return nil, fmt.Errorf("http: %d GET %s", resp.StatusCode, srcURL)
	}

	return resp.Body, nil
}

func do() error {
	r, err := fetch()
	if err != nil {
		return err
	}

	events, err := parseEventsFromHTML(r)
	if err != nil {
		return err
	}
	defer r.Close()

	log.Println("events_count: ", len(events))

	return save(events)
}

func main() {
	if err := do(); err != nil {
		log.Printf("ERROR %s", err)
	}
	for _ = range time.NewTicker(waitTimeout).C {
		if err := do(); err != nil {
			log.Printf("ERROR %s", err)
		}
	}
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}

func notEmptyStr(val, defaultVal string) string {
	if val != "" {
		return val
	}
	return defaultVal
}

func env(key, defaultVal string) string {
	return notEmptyStr(os.Getenv(key), defaultVal)
}
