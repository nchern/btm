package main

import (
	"encoding/json"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustOpenAsset(name string) io.ReadCloser {
	r, err := os.Open(name)
	dieIf(err)
	return r
}

func TestParse(t *testing.T) {

	expectedReader := mustOpenAsset("./testdata/sample.json")
	var expected [][]interface{}
	err := json.NewDecoder(expectedReader).Decode(&expected)
	dieIf(err)

	inputReader := mustOpenAsset("./testdata/sample.html")

	actual, err := parseEventsFromHTML(inputReader)
	assert.NoError(t, err)

	assert.Equal(t, expected, actual)
}
