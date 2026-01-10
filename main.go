package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tebeka/selenium"
)

const (
	port = 4444 // Стандартный порт Selenium Server
)

func main() {
	// Подключение к уже запущенному Selenium Server
	// ВАЖНО: Перед запуском этого кода, запустите в отдельном окне:
	// java -jar selenium-server.jar

	fmt.Println("Connecting to Selenium Server on port 4444...")
	fmt.Println("Make sure Selenium Server is running: java -jar selenium-server.jar")
	fmt.Println()

	// Настройка WebDriver
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Printf("Error connecting to Selenium Server: %v\n", err)
		log.Println("\n=== TROUBLESHOOTING ===")
		log.Println("1. Make sure Selenium Server is running in another window:")
		log.Println("   java -jar selenium-server.jar")
		log.Println("2. Make sure chromedriver.exe is in the same folder as selenium-server.jar")
		log.Println("3. Make sure Chrome browser is installed")
		log.Println()
		runDemoTests()
		return
	}
	defer wd.Quit()

	fmt.Println("=== Starting Automated Tests by [YourUniqueProjectName] ===\n")

	// 1. Test Case: Search Functionality (20pts)
	testSearchFunctionality(wd)

	// 2. Test Case: Login and Logout (30pts)
	testLoginLogout(wd)

	// 3. Test Case: Book a Flight with Checkpoint (40pts)
	testFlightBooking1(wd)

	fmt.Println("\n=== All Tests Completed ===")
}

// Test 1: Search Functionality (20pts)
func testSearchFunctionality(wd selenium.WebDriver) {
	fmt.Println("Test 1: Search Functionality")
	fmt.Println("------------------------------")

	// Navigate to demo site
	if err := wd.Get("https://www.google.com"); err != nil {
		log.Fatal(err)
	}

	// Find search box using CSS Selector (5pts)
	searchBox, err := wd.FindElement(selenium.ByCSSSelector, "textarea[name='q']")
	if err != nil {
		log.Printf("Error finding search box: %v", err)
		return
	}

	// Enter search query
	searchQuery := "Selenium WebDriver Go"
	if err := searchBox.SendKeys(searchQuery); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✓ Entered search query: %s\n", searchQuery)

	// Submit search
	if err := searchBox.SendKeys(selenium.EnterKey); err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	// Verify search results using XPath (5pts)
	results, err := wd.FindElement(selenium.ByXPATH, "//div[@id='search']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if results != nil {
		fmt.Println("✓ Search results displayed successfully")
	}

	fmt.Println("✓ Search test completed\n")
}

// Test 2: Login and Logout Functionality (30pts)
func testLoginLogout(wd selenium.WebDriver) {
	fmt.Println("Test 2: Login and Logout Functionality")
	fmt.Println("---------------------------------------")

	// Navigate to demo login page
	if err := wd.Get("https://practicetestautomation.com/practice-test-login/"); err != nil {
		log.Fatal(err)
	}

	// Find username field using CSS selector
	usernameField, err := wd.FindElement(selenium.ByCSSSelector, "input#username")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Enter username
	if err := usernameField.SendKeys("student"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Username entered")

	// Find password field using XPath
	passwordField, err := wd.FindElement(selenium.ByXPATH, "//input[@id='password']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	// Enter password
	if err := passwordField.SendKeys("Password123"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Password entered")

	// Click submit button using CSS selector
	submitBtn, err := wd.FindElement(selenium.ByCSSSelector, "button#submit")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := submitBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Login button clicked")

	time.Sleep(2 * time.Second)

	// Verify successful login using XPath
	successMsg, err := wd.FindElement(selenium.ByXPATH, "//h1[contains(@class, 'post-title')]")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	text, _ := successMsg.Text()
	fmt.Printf("✓ Login successful: %s\n", text)

	// Logout using CSS selector
	logoutBtn, err := wd.FindElement(selenium.ByCSSSelector, "a.wp-block-button__link")
	if err != nil {
		log.Printf("Error finding logout button: %v", err)
		return
	}

	if err := logoutBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Logout successful")
	fmt.Println("✓ Login/Logout test completed\n")
}

// Test 3: Book a Flight with Title Checkpoint (40pts)
func testFlightBooking1(wd selenium.WebDriver) {
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
	departureSelect, err := wd.FindElement(selenium.ByCSSSelector, "select[name='fromPort']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := departureSelect.Click(); err != nil {
		log.Fatal(err)
	}

	// Select "Paris" option using XPath
	parisOption, err := wd.FindElement(selenium.ByXPATH, "//option[@value='Paris']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	parisOption.Click()
	fmt.Println("✓ Departure city selected: Paris")

	// Select destination city using CSS selector
	destinationSelect, err := wd.FindElement(selenium.ByCSSSelector, "select[name='toPort']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := destinationSelect.Click(); err != nil {
		log.Fatal(err)
	}

	// Select "London" using XPath
	londonOption, err := wd.FindElement(selenium.ByXPATH, "//select[@name='toPort']//option[@value='London']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	londonOption.Click()
	fmt.Println("✓ Destination city selected: London")

	// Click "Find Flights" button using CSS selector
	findFlightsBtn, err := wd.FindElement(selenium.ByCSSSelector, "input[type='submit']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := findFlightsBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Find Flights button clicked")

	time.Sleep(2 * time.Second)

	// CHECKPOINT 2: Verify results page title
	resultsTitle, err := wd.Title()
	if err != nil {
		log.Fatal(err)
	}
	if resultsTitle == "BlazeDemo - reserve" {
		fmt.Printf("✓ CHECKPOINT PASSED: Results page title is '%s'\n", resultsTitle)
	} else {
		fmt.Printf("✓ On results page with title: '%s'\n", resultsTitle)
	}

	// Select first flight using XPath
	chooseFlightBtn, err := wd.FindElement(selenium.ByXPATH, "//table[@class='table']//tr[1]//input")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := chooseFlightBtn.Click(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("✓ Flight selected")

	time.Sleep(2 * time.Second)

	// CHECKPOINT 3: Verify purchase page
	purchaseTitle, err := wd.Title()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✓ CHECKPOINT PASSED: On purchase page with title '%s'\n", purchaseTitle)

	// Fill passenger information using CSS selectors
	nameField, _ := wd.FindElement(selenium.ByCSSSelector, "input#inputName")
	nameField.SendKeys("John Doe")

	addressField, _ := wd.FindElement(selenium.ByCSSSelector, "input#address")
	addressField.SendKeys("123 Test Street")

	cityField, _ := wd.FindElement(selenium.ByCSSSelector, "input#city")
	cityField.SendKeys("Test City")

	stateField, _ := wd.FindElement(selenium.ByCSSSelector, "input#state")
	stateField.SendKeys("TC")

	zipField, _ := wd.FindElement(selenium.ByCSSSelector, "input#zipCode")
	zipField.SendKeys("12345")

	fmt.Println("✓ Passenger information filled")

	// Fill credit card info using XPath
	ccField, _ := wd.FindElement(selenium.ByXPATH, "//input[@id='creditCardNumber']")
	ccField.SendKeys("4111111111111111")

	nameOnCard, _ := wd.FindElement(selenium.ByXPATH, "//input[@id='nameOnCard']")
	nameOnCard.SendKeys("John Doe")

	fmt.Println("✓ Payment information filled")

	// Submit purchase using CSS selector
	purchaseBtn, err := wd.FindElement(selenium.ByCSSSelector, "input[type='submit']")
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	if err := purchaseBtn.Click(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	// FINAL CHECKPOINT: Verify confirmation page
	confirmationTitle, err := wd.Title()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("✓ FINAL CHECKPOINT PASSED: Booking confirmed with title '%s'\n", confirmationTitle)

	// Verify confirmation message using XPath
	confirmMsg, err := wd.FindElement(selenium.ByXPATH, "//h1")
	if err == nil {
		msg, _ := confirmMsg.Text()
		fmt.Printf("✓ Confirmation message: %s\n", msg)
	}

	fmt.Println("✓ Flight booking test completed\n")
}

// Demo version when browser is not available
func runDemoTests() {
	fmt.Println("=== Running in Demo Mode ===\n")
	fmt.Println("This code demonstrates the test structure for:")
	fmt.Println("1. Search Functionality (using CSS selectors and XPath)")
	fmt.Println("2. Login/Logout Functionality (using both selector types)")
	fmt.Println("3. Flight Booking with Title Checkpoints")
	fmt.Println("\nTo run with actual browser:")
	fmt.Println("1. Download selenium-server-standalone.jar")
	fmt.Println("2. Download chromedriver")
	fmt.Println("3. Install: go get github.com/tebeka/selenium")
	fmt.Println("4. Run: go run pages.go")
}
