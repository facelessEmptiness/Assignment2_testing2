package tests

import (
	"testing"

	"github.com/facelessEmptiness/Assignment2_testing/pages"
	"github.com/facelessEmptiness/Assignment2_testing/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoginLogoutFunctionality(t *testing.T) {
	wd, service, err := utils.SetupWebDriver("chrome")
	if err != nil {
		t.Fatalf("Failed to setup WebDriver: %v", err)
	}
	defer wd.Quit()
	defer service.Stop()

	loginPage := &pages.LoginPage{BasePage: pages.BasePage{WebDriver: wd}}

	// Используем тестовый сайт для демонстрации
	t.Run("TestSuccessfulLogin", func(t *testing.T) {
		err := loginPage.NavigateTo("https://the-internet.herokuapp.com/login")
		assert.NoError(t, err)

		// Логин с простыми локаторами
		err = loginPage.Login("tomsmith", "SuperSecretPassword!")
		assert.NoError(t, err)

		// Проверка успешного логина
		isLoggedIn, err := loginPage.IsLoggedIn()
		assert.NoError(t, err)
		assert.True(t, isLoggedIn, "User should be logged in")

		// Сохраняем скриншот
		utils.TakeScreenshot(wd, "login_success.png")
	})

	t.Run("TestLogout", func(t *testing.T) {
		// Логаут с CSS селектором
		err := loginPage.LogoutWithCSS()
		assert.NoError(t, err)

		// Проверка логаута
		isLoggedOut, err := loginPage.IsLoggedOut()
		assert.NoError(t, err)
		assert.True(t, isLoggedOut, "User should be logged out")

		utils.TakeScreenshot(wd, "logout_success.png")
	})

	t.Run("TestFailedLogin", func(t *testing.T) {
		err := loginPage.NavigateTo("https://the-internet.herokuapp.com/login")
		assert.NoError(t, err)

		// Неверные credentials
		err = loginPage.LoginWithXPath("wronguser", "wrongpass")
		assert.NoError(t, err)

		// Проверка сообщения об ошибке
		hasError, err := loginPage.HasErrorMessage()
		assert.NoError(t, err)
		assert.True(t, hasError, "Should show error message for invalid login")
	})
}
