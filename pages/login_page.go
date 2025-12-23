package pages

import (
	"github.com/tebeka/selenium"
)

type LoginPage struct {
	BasePage
}

func (p *LoginPage) NavigateTo(url string) error {
	return p.WebDriver.Get(url)
}

func (p *LoginPage) Login(username, password string) error {
	// Простые локаторы
	err := p.WaitAndSendKeys(selenium.ByID, "username", username)
	if err != nil {
		return err
	}

	err = p.WaitAndSendKeys(selenium.ByID, "password", password)
	if err != nil {
		return err
	}

	return p.WaitAndClick(selenium.ByCSSSelector, "button[type='submit']")
}

func (p *LoginPage) IsLoggedIn() (bool, error) {
	_, err := p.WaitForElement(selenium.ByCSSSelector, ".flash.success")
	return err == nil, err
}

func (p *LoginPage) LogoutWithCSS() error {
	// CSS селектор для кнопки логаута
	return p.WaitAndClick(selenium.ByCSSSelector, "a[href='/logout']")
}

func (p *LoginPage) IsLoggedOut() (bool, error) {
	_, err := p.WaitForElement(selenium.ByID, "username")
	return err == nil, err
}

func (p *LoginPage) LoginWithXPath(username, password string) error {
	// XPath для полей ввода
	err := p.WaitAndSendKeys(selenium.ByXPATH, "//input[@id='username']", username)
	if err != nil {
		return err
	}

	err = p.WaitAndSendKeys(selenium.ByXPATH, "//input[@id='password']", password)
	if err != nil {
		return err
	}

	return p.WaitAndClick(selenium.ByXPATH, "//button[@type='submit']")
}

func (p *LoginPage) HasErrorMessage() (bool, error) {
	_, err := p.WaitForElement(selenium.ByXPATH, "//div[contains(@class, 'flash error')]")
	return err == nil, err
}
