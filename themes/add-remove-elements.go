package themes

import (
	"log"
	"selenium/internal/config"
	"time"

	"github.com/tebeka/selenium"
)

func AddRemoveElements() {
	LogStartTest("Add/Remove Elements")

	time.Sleep(1 * time.Second)

	link, err := config.WD.FindElement(selenium.ByLinkText, "Add/Remove Elements")

	condition := func(wd selenium.WebDriver) (bool, error) {
		elem, err := wd.FindElement(selenium.ByLinkText, "Add/Remove Elements")
		if err != nil {
			return false, nil
		}
		return elem.IsDisplayed()
	}

	err = config.WD.WaitWithTimeout(condition, 5*time.Second)

	if err != nil {
		LogTestError("Timeout", "Ссылка не появилась в течение 10 секунд: %v", err)
		return
	}

	err = link.Click()

	if err != nil {
		LogTestError("Переход по ссылке", "Не удалось перейти по ссылке")
	}

	LogTestInfo("Action", "Успешный переход по ссылке")

	addElementBtn, err := config.WD.FindElement(selenium.ByXPATH, "//button[@onclick='addElement()']")

	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки: %v", err)
	}

	LogTestInfo("Action", "Добавление элементов (5 раз)")
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		err = addElementBtn.Click()
		if err != nil {
			log.Fatalf("Ошибка при клике по кнопке: %v", err)
		}

		LogStep("Кнопка Add Element %sуспешно нажата%s", colorGreen, colorReset)
	}

	deleteButtons, err := config.WD.FindElements(selenium.ByCSSSelector, ".added-manually")

	if err != nil {
		log.Fatalf("Ошибка поиска кнопок удаления: %v", err)
	}

	LogTestInfo("Action", "Удаление элементов (5)")
	for _, el := range deleteButtons {
		time.Sleep(1 * time.Second)
		err := el.Click()
		if err != nil {
			log.Fatalf("Ошибка при нажатии накнопку: %v", err)
		}
		LogStep("Кнопка Delete %sуспешно нажата%s", colorGreen, colorReset)
	}

	deleteButtons, err = config.WD.FindElements(selenium.ByCSSSelector, ".added-manually")

	if len(deleteButtons) == 0 {
		LogTestPassed("Тест Add/Remove Elements", "Успешно завершен.")
	}

	config.WD.Get("https://the-internet.herokuapp.com/")

}
