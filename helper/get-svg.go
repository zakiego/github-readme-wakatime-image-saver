package helper

import (
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
)

func GetSVG(url string) string {
	var body []byte

	retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			if resp.StatusCode != 200 {
				return err
			}

			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		},
	)

	// resp, err := http.Get(url)

	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(resp.StatusCode)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	return string(body)
}
