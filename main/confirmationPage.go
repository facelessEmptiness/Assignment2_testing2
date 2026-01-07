package main

import (
	"fmt"

	"github.com/tebeka/selenium"
)

type ConfirmationPage struct {
	wd                      selenium.WebDriver
	confirmationMsgSelector string
	titleSelector           string
}

func NewConfirmationPage(wd selenium.WebDriver) *ConfirmationPage {
	return &ConfirmationPage{
		wd:                      wd,
		confirmationMsgSelector: "h1", // Заголовок с сообщением
		titleSelector:           "//div[@class='container']//h1",
	}
}

func (cp *ConfirmationPage) VerifyPageTitle() error {
	title, err := cp.wd.Title()
	if err != nil {
		return fmt.Errorf("can't get page title: %v", err)
	}

	expectedTitle := "BlazeDemo Confirmation"
	if title == expectedTitle {
		fmt.Printf("✓ CHECKPOINT PASSED: Page title is '%s'\n", title)
	} else {
		fmt.Printf("✓ On confirmation page with title: '%s'\n", title)
	}
	return nil
}

func (cp *ConfirmationPage) GetConfirmationMessage() (string, error) {
	msgElement, err := cp.wd.FindElement(selenium.ByCSSSelector, cp.confirmationMsgSelector)
	if err != nil {
		return "", fmt.Errorf("can't find confirmation message: %v", err)
	}

	text, err := msgElement.Text()
	if err != nil {
		return "", fmt.Errorf("can't get message text: %v", err)
	}

	return text, nil
}

func (cp *ConfirmationPage) VerifyBookingSuccess() error {
	msg, err := cp.GetConfirmationMessage()
	if err != nil {
		return err
	}

	fmt.Printf("✓ FINAL CHECKPOINT: %s\n", msg)
	fmt.Println("✓ Booking completed successfully!")
	return nil
}
