package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

type Booking struct {
	wd                    selenium.WebDriver
	SelectDepartureCity   string
	SelectDestinationCity string
	findFlightBtn         string
}

func NewBooking(wd selenium.WebDriver) *Booking {
	return &Booking{
		wd:                    wd,
		SelectDepartureCity:   "select[name='fromPort']",
		SelectDestinationCity: "select[name='toPort']",
		findFlightBtn:         "input[type='submit']",
	}

}
func (b *Booking) SelectDepartureCityBtn(city string) error {
	departureSelect, err := b.wd.FindElement(selenium.ByCSSSelector, b.SelectDepartureCity)

	if err != nil {
		fmt.Printf("Can't find departure dropdown: %v", err)
	}

	if err := departureSelect.Click(); err != nil {
		return err
	}

	option, err := b.wd.FindElement(selenium.ByXPATH,
		fmt.Sprintf("//select[@name='fromPort']//option[@value='%s']", city))

	if err != nil {
		return fmt.Errorf("can't find city %s: %v", city, err)
	}

	if err := option.Click(); err != nil {
		return err
	}

	fmt.Printf("✓ Departure city selected: %s\n", city)
	return nil
}
func (b *Booking) SelectDestinationCityBtn(city string) error {
	destinationSelect, err := b.wd.FindElement(selenium.ByCSSSelector, b.SelectDestinationCity)

	if err != nil {
		fmt.Printf("Can't find destination dropdown: %v", err)
	}

	if err := destinationSelect.Click(); err != nil {
		return err
	}

	option, err := b.wd.FindElement(selenium.ByXPATH, fmt.Sprintf("//select[@name='toPort']//option[@value='%s']", city))
	if err != nil {
		return fmt.Errorf("can't find destination dropdown %s: %v", city, err)
	}

	if err := option.Click(); err != nil {
		return err
	}
	fmt.Printf("Destination city selected: %s\n", city)
	return nil
}
func (b *Booking) clickFindFlights() error {
	findFlightsBtn, err := b.wd.FindElement(selenium.ByCSSSelector, b.findFlightBtn)
	if err != nil {
		log.Printf("Can't find button: %v", err)

	}
	if err := findFlightsBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Find Flights button clicked")
	time.Sleep(2 * time.Second)
	return nil
}

// Test 3: Book a Flight with Title Checkpoint (40pts)
func testFlightBooking(wd selenium.WebDriver) {
	fmt.Println("Test 3: Flight Booking with Checkpoint")
	fmt.Println("----------------------------------------")

	// Navigate to demo flight booking site
	if err := wd.Get("https://blazedemo.com/"); err != nil {
		log.Fatal(err)
	}

	// CHECKPOINT 1: Verify page title
	title, err := wd.Title()
	if err != nil {
		log.Fatal(err)
	}
	expectedTitle := "BlazeDemo"
	if title == expectedTitle {
		fmt.Printf("✓ CHECKPOINT PASSED: Page title is '%s'\n", title)
	} else {
		fmt.Printf("✗ CHECKPOINT FAILED: Expected '%s', got '%s'\n", expectedTitle, title)
	}

	//Select departure city using CSS selector
	BookingPage := NewBooking(wd)
	if err := BookingPage.SelectDepartureCityBtn("Paris"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := BookingPage.SelectDestinationCityBtn("London"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := BookingPage.clickFindFlights(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("✓ Test completed successfully!")
}
