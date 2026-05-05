package main

import (
	"selenium/internal/config"
	"selenium/themes"
	"time"
)

func main() {
	config.InitSelenium()

	defer config.WD.Quit()

	config.WD.Get("https://the-internet.herokuapp.com/")

	themes.AddRemoveElements()

	time.Sleep(4 * time.Second)

}
