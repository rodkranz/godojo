package kata_01

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"regexp"
	"fmt"
)

func TestRequest_CannotMakeARequestWithInvalidURL(t *testing.T) {
	_, err := Request("full")
	assert.EqualError(t, err, "Invalid URL" )
}

func TestRequest_ShouldReturnNilForAValidURL(t *testing.T) {
	_, err := Request("http://www.google.com")

	assert.Equal(t, nil, err, "Url must return valid.")
}

func TestRequest_ShouldReturnTimeFormat(t *testing.T) {
	resp, _ := Request("http://www.google.com")
	assertValidUrl(t, resp);
}

func TestRun_ShouldMakeMoreThanOneRequest(t *testing.T) {
	urls := []string{"http://www.google.com", "http://www.google.com"}
	outputs, _ := Run(urls)

	for _, output := range outputs {
		assertValidUrl(t, output)
	}

}

func assertValidUrl (t *testing.T, output string) {
	eg := regexp.MustCompile(`^\[([\d,.]+s\])+([\w:\ ]+)(.*)(with\ )+(\[+\d+\])`)
	assert.True(t, eg.Match([]byte(output)), "Must return with time")
}