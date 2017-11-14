package kata_01

import "errors"
import (
	"net/url"
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
)

func Run(urls []string) ([]string, error) {
	startTime := time.Now()
	results := make([]string, 0)

	var err error
	for _, url := range urls {
		var output string
		output, err = Request(url)
		results = append(results, output)
	}

	endTime := time.Since(startTime)
	results = append(results, fmt.Sprintf("[%.2fs] elapsed time.", endTime.Seconds()))
	return results, err
}

// emotion a lot
func Request(requestUrl string) (string, error) {
	validUrl, err := url.ParseRequestURI(requestUrl)
	if err != nil {
		return "", errors.New("Invalid URL")
	}

	startTime := time.Now()
	response, requestError := http.Get(validUrl.String())

	if requestError != nil {
		return "", requestError;
	}

	endTime := time.Since(startTime)

	//[0.98s] elapsed time for request [http://www.terra.com.br/] with [262425] bytes

	byte, _ := ioutil.ReadAll(response.Body)

	return fmt.Sprintf("[%.2fs] elapsed time for request [%s] with [%d] bytes",
		endTime.Seconds(), validUrl.String(), len(byte)), nil
}
