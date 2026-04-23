package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func APIWakatimeStats(config Config) string {

	wakatime := config.Wakatime

	host := "https://github-readme-stats.vercel.app/api/wakatime"
	API := fmt.Sprintf("%s?username=%s&hide_border=%s&langs_count=%s&bg_color=00000000", host, wakatime.Username, strconv.FormatBool(wakatime.HideBorder), strconv.Itoa(wakatime.LangsCount))
	return API
}

func APIGithubStats(config Config) string {

	github := config.Github

	host := "https://github-readme-stats.vercel.app/api?"

	var query = func() string {
		username := "username=" + github.Username
		show_icons := "show_icons=" + strconv.FormatBool(github.ShowIcons)
		hide_border := "hide_border=" + strconv.FormatBool(github.HideBorder)
		count_private := "count_private=" + strconv.FormatBool(github.CountPrivate)
		bg_color := "bg_color=00000000"

		return strings.Join([]string{username, show_icons, hide_border, count_private, bg_color}, "&")
	}

	API := host + query()
	return API
}
