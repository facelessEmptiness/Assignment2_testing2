package pages

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type PurchasePage struct {
	wd                  selenium.WebDriver
	nameSelector        string
	addressSelector     string
	citySelector        string
	stateSelector       string
	zipCodeSelector     string
	ccNumberSelector    string
	nameOnCardSelector  string
	purchaseBtnSelector string
}

func NewPurchasePage(wd selenium.WebDriver) *PurchasePage {
	return &PurchasePage{
		wd:                  wd,
		nameSelector:        "input#inputName",
		addressSelector:     "input#address",
		citySelector:        "input#city",
		stateSelector:       "input#state",
		zipCodeSelector:     "input#zipCode",
		ccNumberSelector:    "input#creditCardNumber",
		nameOnCardSelector:  "input#nameOnCard",
		purchaseBtnSelector: "input[type='submit']",
	}
}

func (p *PurchasePage) FillPassengerInfo(name, address, city, state, zipCode string) error {
	// Имя
	nameField, err := p.wd.FindElement(selenium.ByCSSSelector, p.nameSelector)
	if err != nil {
		return fmt.Errorf("can't find name field: %v", err)
	}
	if err := nameField.SendKeys(name); err != nil {
		return err
	}

	// Адрес
	addressField, err := p.wd.FindElement(selenium.ByCSSSelector, p.addressSelector)
	if err != nil {
		return fmt.Errorf("can't find address field: %v", err)
	}
	if err := addressField.SendKeys(address); err != nil {
		return err
	}

	// Город
	cityField, err := p.wd.FindElement(selenium.ByCSSSelector, p.citySelector)
	if err != nil {
		return fmt.Errorf("can't find city field: %v", err)
	}
	if err := cityField.SendKeys(city); err != nil {
		return err
	}

	// Штат
	stateField, err := p.wd.FindElement(selenium.ByCSSSelector, p.stateSelector)
	if err != nil {
		return fmt.Errorf("can't find state field: %v", err)
	}
	if err := stateField.SendKeys(state); err != nil {
		return err
	}

	// Индекс
	zipField, err := p.wd.FindElement(selenium.ByCSSSelector, p.zipCodeSelector)
	if err != nil {
		return fmt.Errorf("can't find zip code field: %v", err)
	}
	if err := zipField.SendKeys(zipCode); err != nil {
		return err
	}

	fmt.Println("✓ Passenger information filled")
	return nil
}

func (p *PurchasePage) FillPaymentInfo(ccNumber, nameOnCard string) error {
	// Номер карты
	ccField, err := p.wd.FindElement(selenium.ByCSSSelector, p.ccNumberSelector)
	if err != nil {
		return fmt.Errorf("can't find credit card field: %v", err)
	}
	if err := ccField.SendKeys(ccNumber); err != nil {
		return err
	}

	// Имя на карте
	nameField, err := p.wd.FindElement(selenium.ByCSSSelector, p.nameOnCardSelector)
	if err != nil {
		return fmt.Errorf("can't find name on card field: %v", err)
	}
	if err := nameField.SendKeys(nameOnCard); err != nil {
		return err
	}

	fmt.Println("✓ Payment information filled")
	return nil
}

func (p *PurchasePage) ClickPurchase() error {
	btn, err := p.wd.FindElement(selenium.ByCSSSelector, p.purchaseBtnSelector)
	if err != nil {
		return fmt.Errorf("can't find purchase button: %v", err)
	}

	if err := btn.Click(); err != nil {
		return fmt.Errorf("can't click purchase button: %v", err)
	}

	fmt.Println("✓ Purchase button clicked")
	return nil
}
