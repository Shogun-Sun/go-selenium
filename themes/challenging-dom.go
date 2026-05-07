package themes

import (
	"fmt"
	"regexp"
	"selenium/internal/config"

	"github.com/tebeka/selenium"
)

func ChallengingDom() {
	RunTest("Challenging Dom", func() {
		config.WD.Get("https://the-internet.herokuapp.com/challenging_dom")

		btnBlue, err := config.WD.FindElement(selenium.ByXPATH, "//a[@class='button']")
		if err != nil {
			LogTestError("Ошибка", "Не удалось найти синюю кнопк: %v", err)
			return
		}
		btnRed, err := config.WD.FindElement(selenium.ByXPATH, "//a[@class='button alert']")
		if err != nil {
			LogTestError("Ошибка", "Не удалось найти красную кнопку")
			return
		}
		btnGreen, err := config.WD.FindElement(selenium.ByXPATH, "//a[@class='button success']")
		if err != nil {
			LogTestError("Ошибка", "Не удалось найти зеленую кнопку")
			return
		}

		LogTestInfo("Action", "Все кнопки найдены")

		err = btnBlue.Click()

		err = btnRed.Click()

		err = btnGreen.Click()

		// ============== Получение вопроса =================
		source, err := config.WD.PageSource()

		if err != nil {
			LogTestError("Ошибка", "Не удалось получить страницу: %v", err)
		}

		re := regexp.MustCompile(`Answer: (\d+)`)
		match := re.FindStringSubmatch(source)

		if len(match) > 1 {
			fmt.Println("Answer:", match[1])
		}

		// ============== Получение вопроса =================

	})
}
