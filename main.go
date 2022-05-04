package main

import (
	"github.com/zakiego/github-readme-wakatime-image-saver/helper"
)

func main() {
	config := helper.ReadConfig()

	APIGithub := helper.APIGithubStats(config)
	SVGGithub := helper.GetSVG(APIGithub)
	helper.SaveSVG(SVGGithub, "github-stats.svg")

	APIWakatime := helper.APIWakatimeStats(config)
	SVGWakatime := helper.GetSVG(APIWakatime)
	helper.SaveSVG(SVGWakatime, "wakatime-stats.svg")
}
