package pages

import (
	"testing"

	"github.com/facelessEmptiness/Assignment2_testing/utils"
	"github.com/stretchr/testify/assert"
)

func TestSearchFunctionality(t *testing.T) {
	wd, service, err := utils.SetupWebDriver("chrome")
	if err != nil {
		t.Fatalf("Failed to setup WebDriver: %v", err)
	}
	defer wd.Quit()
	defer service.Stop()

	searchPage := &pages.SearchPage{BasePage: pages.BasePage{WebDriver: wd}}

	// Тест 1: Использование простых локаторов
	t.Run("TestSearchWithSimpleLocators", func(t *testing.T) {
		err := searchPage.NavigateTo("https://www.google.com")
		assert.NoError(t, err)

		// By name (простой локатор)
		err = searchPage.EnterSearchQuery("Selenium WebDriver")
		assert.NoError(t, err)

		// By class name (простой локатор)
		err = searchPage.ClickSearchButton()
		assert.NoError(t, err)

		// Проверка результатов
		hasResults, err := searchPage.HasSearchResults()
		assert.NoError(t, err)
		assert.True(t, hasResults, "Search should return results")
	})

	// Тест 2: Использование CSS селектора (5 баллов)
	t.Run("TestSearchWithCSSSelector", func(t *testing.T) {
		err := searchPage.NavigateTo("https://duckduckgo.com")
		assert.NoError(t, err)

		// CSS селектор
		err = searchPage.SearchWithCSS("CSS Selector Test")
		assert.NoError(t, err)

		results, err := searchPage.GetCSSSearchResults()
		assert.NoError(t, err)
		assert.Greater(t, len(results), 0, "Should have search results")
	})

	// Тест 3: Использование XPath (5 баллов)
	t.Run("TestSearchWithXPath", func(t *testing.T) {
		err := searchPage.NavigateTo("https://www.bing.com")
		assert.NoError(t, err)

		// XPath
		err = searchPage.SearchWithXPath("XPath Test")
		assert.NoError(t, err)

		results, err := searchPage.GetXPathSearchResults()
		assert.NoError(t, err)
		assert.Greater(t, len(results), 0, "Should have search results")
	})
}
