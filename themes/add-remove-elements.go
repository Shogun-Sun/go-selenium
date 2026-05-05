package themes

import (
	"fmt"
	"log"
	"selenium/internal/config"

	"github.com/tebeka/selenium"
)

func AddRemoveElements() {

	link, err := config.WD.FindElement(selenium.ByLinkText, "Add/Remove Elements")

	if err != nil {
		log.Fatalf("Ссылка не найдена: %v", err)
	}

	err = link.Click()

	if err != nil {
		log.Fatalf("Ошибка при клике на ссылку: %v", err)
	}

	addElementBtn, err := config.WD.FindElement(selenium.ByXPATH, "//button[@onclick='addElement()']")

	if err != nil {
		log.Fatalf("Ошибка при поиске кнопки: %v", err)
	}

	for i := 0; i < 5; i++ {
		err = addElementBtn.Click()
		if err != nil {
			log.Fatalf("Ошибка при клике по кнопке: %v", err)
		}
	}

	deleteButtons, err := config.WD.FindElements(selenium.ByCSSSelector, ".added-manually")

	if err != nil {
		log.Fatalf("Ошибка поиска кнопок удаления: %v", err)
	}

	for _, el := range deleteButtons {
		elementText, err := el.Text()
		elementType, err := el.TagName()

		if err != nil {
			log.Fatalf("Ошибка получения текста или типа элемента: %v", err)
		}
		fmt.Printf("Найден элемент: %s, текст элемента: %s\n", elementType, elementText)
	}

	for _, el := range deleteButtons {
		err := el.Click()

		if err != nil {
			log.Fatalf("Ошибка при нажатии накнопку: %v", err)
		}
	}

	deleteButtons, err = config.WD.FindElements(selenium.ByCSSSelector, ".added-manually")

	if len(deleteButtons) == 0 {
		log.Print("Все кнопки удалены!")
	}

}
