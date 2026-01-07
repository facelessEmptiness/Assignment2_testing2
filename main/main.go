package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	port = 4444
)

func main() {
	fmt.Println("=== Starting Flight Booking Test with POM ===\n")

	// Настройка WebDriver
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Failed to connect to Selenium Server: %v\n", err)
		log.Println("Make sure Selenium Server is running: java -jar selenium-server.jar")
		return
	}
	defer wd.Quit() // ✅ Закрываем браузер в конце

	fmt.Println("✓ Connected to Selenium Server")
	fmt.Println("Test 3: Flight Booking with Checkpoint")
	fmt.Println("----------------------------------------\n")

	// Navigate to demo flight booking site
	if err := wd.Get("https://blazedemo.com/"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Navigated to BlazeDemo\n")

	// ============ СТРАНИЦА 1: Home Page ============
	fmt.Println("=== Page 1: Home Page ===")
	homePage := NewBooking(wd)

	if err := homePage.SelectDepartureCityBtn("Paris"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := homePage.SelectDestinationCityBtn("London"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := homePage.clickFindFlights(); err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println()

	// ============ СТРАНИЦА 2: Reserve Page ============
	fmt.Println("=== Page 2: Reserve Page ===")
	reservePage := NewChooseFlightBookingPage(wd)

	if err := reservePage.VerifyPageTitle(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := reservePage.SelectFlight(); err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println()

	// ============ СТРАНИЦА 3: Purchase Page ============
	fmt.Println("=== Page 3: Purchase Page ===")
	purchasePage := NewPurchasePage(wd)

	if err := purchasePage.FillPassengerInfo("John Doe", "123 Test St", "Test City", "TC", "12345"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := purchasePage.FillPaymentInfo("4111111111111111", "John Doe"); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := purchasePage.ClickPurchase(); err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Println()

	// Небольшая задержка для загрузки страницы
	time.Sleep(2 * time.Second)

	// ============ СТРАНИЦА 4: Confirmation Page ============
	fmt.Println("=== Page 4: Confirmation Page ===")
	confirmationPage := NewConfirmationPage(wd)

	if err := confirmationPage.VerifyPageTitle(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := confirmationPage.VerifyBookingSuccess(); err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println("✓✓✓ All tests completed successfully! ✓✓✓")
}
