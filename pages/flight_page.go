package pages

import (
	"strings"

	"github.com/tebeka/selenium"
)

type FlightPage struct {
	BasePage
}

func (p *FlightPage) NavigateTo(url string) error {
	return p.WebDriver.Get(url)
}

func (p *FlightPage) GetPageTitle() (string, error) {
	// Проверка заголовка страницы (checkpoint)
	return p.WebDriver.Title()
}

func (p *FlightPage) SelectDepartureCity(city string) error {
	// Простой локатор - By name
	return p.SelectFromDropdown(selenium.ByName, "fromPort", city)
}

func (p *FlightPage) SelectDestinationCity(city string) error {
	// Простой локатор - By name
	return p.SelectFromDropdown(selenium.ByName, "toPort", city)
}

func (p *FlightPage) SelectFromDropdown(by, selector, value string) error {
	element, err := p.WaitForElement(by, selector)
	if err != nil {
		return err
	}

	selectElem, err := element.FindElement(selenium.ByXPATH, ".//option[text()='"+value+"']")
	if err != nil {
		return err
	}

	return selectElem.Click()
}

func (p *FlightPage) FindFlights() error {
	// CSS селектор для кнопки поиска
	return p.WaitAndClick(selenium.ByCSSSelector, "input[type='submit']")
}

func (p *FlightPage) HasAvailableFlights() (bool, error) {
	elements, err := p.WebDriver.FindElements(selenium.ByCSSSelector, "table.table tbody tr")
	return len(elements) > 0, err
}

func (p *FlightPage) ChooseFirstFlight() error {
	// XPath для выбора первого рейса
	return p.WaitAndClick(selenium.ByXPATH, "(//input[@type='submit'])[1]")
}

func (p *FlightPage) FillPassengerInfo(firstName, lastName, email, phone, address string) error {
	// Заполнение формы с использованием разных типов локаторов

	// By ID
	err := p.WaitAndSendKeys(selenium.ByID, "inputName", firstName+" "+lastName)
	if err != nil {
		return err
	}

	// By name
	err = p.WaitAndSendKeys(selenium.ByName, "address", address)
	if err != nil {
		return err
	}

	// CSS селектор
	err = p.WaitAndSendKeys(selenium.ByCSSSelector, "#city", "New York")
	if err != nil {
		return err
	}

	err = p.WaitAndSendKeys(selenium.ByCSSSelector, "#state", "NY")
	if err != nil {
		return err
	}

	// XPath
	err = p.WaitAndSendKeys(selenium.ByXPATH, "//input[@id='zipCode']", "10001")
	if err != nil {
		return err
	}

	// Выбор типа карты с CSS селектором
	err = p.WaitAndClick(selenium.ByCSSSelector, "select#cardType option[value='visa']")
	if err != nil {
		return err
	}

	// Заполнение данных карты
	err = p.WaitAndSendKeys(selenium.ByID, "creditCardNumber", "4111111111111111")
	if err != nil {
		return err
	}

	err = p.WaitAndSendKeys(selenium.ByID, "creditCardMonth", "12")
	if err != nil {
		return err
	}

	err = p.WaitAndSendKeys(selenium.ByID, "creditCardYear", "2025")
	if err != nil {
		return err
	}

	return p.WaitAndSendKeys(selenium.ByID, "nameOnCard", firstName+" "+lastName)
}

func (p *FlightPage) GetTotalPrice() (string, error) {
	element, err := p.WaitForElement(selenium.ByXPATH, "//p[contains(text(), 'Total Cost')]")
	if err != nil {
		return "", err
	}

	text, err := element.Text()
	if err != nil {
		return "", err
	}

	// Извлекаем только цену
	parts := strings.Split(text, ":")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[1]), nil
	}

	return text, nil
}

func (p *FlightPage) PurchaseFlight() error {
	return p.WaitAndClick(selenium.ByCSSSelector, "input[type='submit']")
}

func (p *FlightPage) IsBookingConfirmed() (bool, error) {
	_, err := p.WaitForElement(selenium.ByXPATH, "//h1[contains(text(), 'Thank you')]")
	return err == nil, err
}

func (p *FlightPage) GetBookingID() (string, error) {
	element, err := p.WaitForElement(selenium.ByXPATH, "//td[text()='Id']/following-sibling::td")
	if err != nil {
		return "", err
	}
	return element.Text()
}
