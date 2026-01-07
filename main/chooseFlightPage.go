package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

type ChooseFlightBookingPage struct {
	wd                     selenium.WebDriver
	ChooseFlightBookingBtn string
}

func NewChooseFlightBookingPage(wd selenium.WebDriver) *ChooseFlightBookingPage {
	return &ChooseFlightBookingPage{
		wd:                     wd,
		ChooseFlightBookingBtn: "//table[@class='table']//tr[1]//input",
	}
}
func (p *ChooseFlightBookingPage) SelectFlight() error {
	chooseFlightBtn, err := p.wd.FindElement(selenium.ByXPATH, p.ChooseFlightBookingBtn)
	if err != nil {
		log.Printf("Can't find choose flight button: %v", err)
		return err
	}

	if err := chooseFlightBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Flight selected")

	time.Sleep(2 * time.Second)
	return nil
}
func (p *ChooseFlightBookingPage) VerifyPageTitle() error {
	title, err := p.wd.Title()
	if err != nil {
		return fmt.Errorf("can't get title: %v", err)
	}

	expectedTitle := "BlazeDemo - reserve"
	if title == expectedTitle {
		fmt.Printf("✓ CHECKPOINT PASSED: Page title is '%s'\n", title)
	} else {
		fmt.Printf("⚠ Warning: Expected '%s', got '%s'\n", expectedTitle, title)
	}
	return nil
}
func ChooseFlightBooking(wd selenium.WebDriver) {
	page := NewChooseFlightBookingPage(wd)
	if err := page.SelectFlight(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

}
