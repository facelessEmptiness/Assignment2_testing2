package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func SetupWebDriver(browser string) (selenium.WebDriver, *selenium.Service, error) {
	// Запускаем Selenium server
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444, opts...)
	if err != nil {
		return nil, nil, err
	}

	// Настройка WebDriver
	caps := selenium.Capabilities{"browserName": browser}

	if browser == "chrome" {
		caps.AddChrome(selenium.Capabilities{
			"args": []string{
				"--no-sandbox",
				"--disable-dev-shm-usage",
				"--disable-gpu",
				"--window-size=1920,1080",
			},
		})
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 4444))
	if err != nil {
		service.Stop()
		return nil, nil, err
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)
	return wd, service, nil
}

func TakeScreenshot(wd selenium.WebDriver, filename string) {
	screenshot, err := wd.Screenshot()
	if err != nil {
		log.Printf("Failed to take screenshot: %v", err)
		return
	}

	err = os.WriteFile(filename, screenshot, 0644)
	if err != nil {
		log.Printf("Failed to save screenshot: %v", err)
	}
}
