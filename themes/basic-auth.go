package themes

import (
	"fmt"
	"selenium/internal/config"
	"time"

	"github.com/tebeka/selenium"
)

func BasicAuth() {
	LogTestInfo("Тест Basic Auth", "Начат.")

	time.Sleep(1 * time.Second)

	link, err := config.WD.FindElement(selenium.ByLinkText, "Basic Auth")

	if err != nil {
		LogTestError("Ошибка", "Ссылка не найдена")
		return
	}

	url, err := link.GetAttribute("href")

	if err != nil {
		LogTestError("Ошибка", "Не удалось получить атрибут href")
		return
	}

	url = fmt.Sprintf("https://the-internet.herokuapp.com%s", url)

	LogTestInfo("Action", "Получена ссылка: %s", url)

	url = fmt.Sprintf("https://%s:%s@the-internet.herokuapp.com/basic_auth", "admin", "admin")

	LogTestInfo("Action", "Итоговая ссылка на авторизацию: %s", url)

	err = config.WD.Get(url)

	if err != nil {
		LogTestError("Ошибка", "Не удалось пройти авторизацию: %v", err)
		return
	}

	config.WD.Get("https://the-internet.herokuapp.com/basic_auth")

	element, err := config.WD.FindElement(selenium.ByTagName, "body")

	if err != nil {
		LogTestError("Ошибка", "Тег body не найден")
		return
	}

	text, err := element.Text()

	if err != nil {
		LogTestError("Ошибка", "Не удалось получить текст")
		return
	}

	if text == "Not authorized" {
		LogTestError("Ошибка", "Авторизация не пройдена")
		return
	}

	element, err = config.WD.FindElement(selenium.ByTagName, "p")

	if err != nil {
		LogTestError("Ошибка", "Тег p не найден")
		return
	}

	text, err = element.Text()

	if err != nil {
		LogTestError("Ошибка", "Не удалось получить текст")
		return
	}

	LogTestInfo("Action", "Результат авторизации: %s", text)
	LogTestPassed("Тест Basic Auth", "Авторизация успешно пройдена")

}
