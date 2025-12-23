package pages

import (
	"github.com/tebeka/selenium"
)

type SearchPage struct {
	BasePage
}

func (p *SearchPage) NavigateTo(url string) error {
	return p.WebDriver.Get(url)
}

func (p *SearchPage) EnterSearchQuery(query string) error {
	// Простой локатор - By name
	return p.WaitAndSendKeys(selenium.ByName, "q", query)
}

func (p *SearchPage) ClickSearchButton() error {
	// Простой локатор - By class name
	return p.WaitAndClick(selenium.ByClassName, "gNO89b")
}

func (p *SearchPage) HasSearchResults() (bool, error) {
	_, err := p.WaitForElement(selenium.ByID, "search")
	return err == nil, err
}

// Методы с CSS селекторами (5 баллов)
func (p *SearchPage) SearchWithCSS(query string) error {
	// CSS селектор
	return p.WaitAndSendKeys(selenium.ByCSSSelector, "input[aria-label='Search']", query)
}

func (p *SearchPage) GetCSSSearchResults() ([]string, error) {
	elements, err := p.WebDriver.FindElements(selenium.ByCSSSelector, ".result__title")
	if err != nil {
		return nil, err
	}

	var results []string
	for _, element := range elements {
		text, _ := element.Text()
		results = append(results, text)
	}
	return results, nil
}

// Методы с XPath (5 баллов)
func (p *SearchPage) SearchWithXPath(query string) error {
	// XPath
	return p.WaitAndSendKeys(selenium.ByXPATH, "//input[@name='q']", query)
}

func (p *SearchPage) GetXPathSearchResults() ([]string, error) {
	elements, err := p.WebDriver.FindElements(selenium.ByXPATH, "//li[@class='b_algo']//h2")
	if err != nil {
		return nil, err
	}

	var results []string
	for _, element := range elements {
		text, _ := element.Text()
		results = append(results, text)
	}
	return results, nil
}
