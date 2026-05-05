package config

import (
	"log"

	"github.com/tebeka/selenium"
)

var WD selenium.WebDriver

func InitSelenium() {
	var err error
	caps := selenium.Capabilities{"browserName": "firefox"}
	WD, err = selenium.NewRemote(caps, "http://localhost:4444/wd/hub")

	if err != nil {
		log.Fatalf("Ошибка при инициализации селениума: %v", err)
	}
}
