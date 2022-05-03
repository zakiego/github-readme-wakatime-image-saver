package helper

import (
	"fmt"
	"strconv"
)

func QueryBuilder(config Config) string {
	host := "https://zakiego-github-readme.vercel.app/api/wakatime"
	API := fmt.Sprintf("%s?username=%s&hide_border=%s&langs_count=%s", host, config.Username, strconv.FormatBool(config.HideBorder), strconv.Itoa(config.LangsCount))
	return API
}
