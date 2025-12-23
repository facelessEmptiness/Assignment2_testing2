package main

import (
	"flag"
	"github.com/facelessEmptiness/a/tests"
	"log"
	"os"
	"os/exec"
	"testing"
)

func main() {
	var (
		startSelenium = flag.Bool("start-selenium", false, "Start Selenium Docker containers")
		testType      = flag.String("test", "all", "Test to run: search, login, flight, or all")
	)
	flag.Parse()

	// Запуск Selenium контейнеров если нужно
	if *startSelenium {
		log.Println("Starting Selenium Docker containers...")
		cmd := exec.Command("docker-compose", "up", "-d")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Failed to start Selenium: %v", err)
		}
		log.Println("Selenium containers started successfully")
	}

	// Запуск тестов
	log.Printf("Running %s tests...", *testType)

	switch *testType {
	case "search":
		m := testing.MainStart(nil, []testing.InternalTest{
			{Name: "TestSearch", F: tests.TestSearchFunctionality},
		}, nil, nil, nil)
		os.Exit(m.Run())
	case "login":
		m := testing.MainStart(nil, []testing.InternalTest{
			{Name: "TestLogin", F: tests.TestLoginLogoutFunctionality},
		}, nil, nil, nil)
		os.Exit(m.Run())
	case "flight":
		m := testing.MainStart(nil, []testing.InternalTest{
			{Name: "TestFlight", F: tests.TestFlightBooking},
		}, nil, nil, nil)
		os.Exit(m.Run())
	case "all":
		m := testing.MainStart(nil, []testing.InternalTest{
			{Name: "TestSearch", F: tests.TestSearchFunctionality},
			{Name: "TestLogin", F: tests.TestLoginLogoutFunctionality},
			{Name: "TestFlight", F: tests.TestFlightBooking},
		}, nil, nil, nil)
		os.Exit(m.Run())
	default:
		log.Fatal("Invalid test type. Use: search, login, flight, or all")
	}
}
