package main

import "github.com/zakiego/gitbuh-readme-action/helper"

func main() {
	config := helper.ReadConfig()

	SVG := helper.GetSVG(helper.QueryBuilder(config))

	helper.SaveSVG(SVG, "stat.svg")
}
