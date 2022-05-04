package helper

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func GetSVG(url string) string {

	c := retryablehttp.NewClient()

	r, err := retryablehttp.NewRequest("GET", url, nil)

	handlerShouldRetry := true

	r.SetResponseHandler(func(*http.Response) error {
		if !handlerShouldRetry {
			return nil
		}
		handlerShouldRetry = false
		return errors.New("retryable error")
	})

	resp, err := c.Do(r)

	fmt.Println("URL :", url)
	fmt.Println("Status Code :", resp.StatusCode)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
