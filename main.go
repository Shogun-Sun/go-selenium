package main

import (
	"selenium/internal/config"
	"selenium/themes"
)

func main() {
	config.InitSelenium()

	defer config.WD.Quit()

	config.WD.Get("https://the-internet.herokuapp.com/")

	themes.AddRemoveElements()
}
