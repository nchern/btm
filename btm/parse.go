package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	eventsVar = regexp.MustCompile(`(?s)\s+boatEvents = \[.*?\];`)

	jsComments = regexp.MustCompile(`(?s)/\*.*?\*/`)
)

func parseEventsFromHTML(r io.Reader) ([][]interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := string(b)
	s = jsComments.ReplaceAllString(s, "")

	for _, s := range eventsVar.FindAllString(s, -1) {
		toks := strings.SplitN(s, "=", 2)
		s = strings.Trim(toks[1], "; \t")

		var events [][]interface{}
		if err := json.Unmarshal([]byte(s), &events); err != nil {
			return nil, err
		}
		return events, nil

	}
	return nil, nil
}
