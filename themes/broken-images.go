package themes

import (
	"fmt"
	"net/http"
	"selenium/internal/config"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func BrokenImages() {
	RunTest("Broken Images", func() {
		config.WD.Get("https://the-internet.herokuapp.com/broken_images")

		elements, err := config.WD.FindElements(selenium.ByTagName, "img")

		if err != nil {
			LogTestError("Ошибка", "Не найжено элементов с тегом img")
		}

		LogTestInfo("Action", "Найдены изображения")

		for _, el := range elements {
			time.Sleep(1 * time.Second)
			srcAttribute, _ := el.GetAttribute("src")
			url := fmt.Sprintf("https://the-internet.herokuapp.com/%s", strings.TrimLeft(srcAttribute, "/"))

			resp, _ := http.Get(url)
			LogStep(url)

			if resp.StatusCode != 200 {
				LogStep("%s%s%s", colorRed, resp.Status, colorReset)
			} else {
				LogStep("%s%s%s", colorGreen, resp.Status, colorReset)
			}
		}

		config.WD.Get("https://the-internet.herokuapp.com")

	})
}
