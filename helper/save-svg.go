package helper

import (
	"fmt"
	"log"
	"os"
)

func SaveSVG(svg string, namefile string) {
	f, err := os.Create(namefile)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	_, err = f.WriteString(svg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("SVG saved")
}
