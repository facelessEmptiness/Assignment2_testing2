package pages

import (
	"time"

	"github.com/tebeka/selenium"
)

type BasePage struct {
	WebDriver selenium.WebDriver
}

func (p *BasePage) WaitForElement(by, value string) (selenium.WebElement, error) {
	return p.WebDriver.FindElement(by, value)
}

func (p *BasePage) WaitAndClick(by, value string) error {
	element, err := p.WaitForElement(by, value)
	if err != nil {
		return err
	}
	return element.Click()
}

func (p *BasePage) WaitAndSendKeys(by, value, text string) error {
	element, err := p.WaitForElement(by, value)
	if err != nil {
		return err
	}
	return element.SendKeys(text)
}

func (p *BasePage) GetText(by, value string) (string, error) {
	element, err := p.WaitForElement(by, value)
	if err != nil {
		return "", err
	}
	return element.Text()
}

func (p *BasePage) Sleep(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
