package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dailymotion/allure-go" // ← Добавь это
	config "github.com/facelessEmptiness/Assignment2_testing2/logs"
	"github.com/facelessEmptiness/Assignment2_testing2/pages"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
)

type BookingTestSuite struct {
	suite.Suite
	wd selenium.WebDriver
}

func (s *BookingTestSuite) takeScreenshot(name string) []byte {
	log.WithField("screenshot", name).Info("Taking screenshot")

	screenshot, err := s.wd.Screenshot()
	if err != nil {
		log.WithError(err).Error("Failed to take screenshot")
		return nil
	}

	// Сохранить в папку screenshots (для резервной копии)
	os.MkdirAll("screenshots", 0755)
	filename := fmt.Sprintf("screenshots/%s_%s.png",
		name,
		time.Now().Format("20060102_150405"),
	)

	err = os.WriteFile(filename, screenshot, 0644)
	if err != nil {
		log.WithError(err).Error("Failed to save screenshot")
	} else {
		log.WithField("file", filename).Info("Screenshot saved")
	}

	return screenshot
}

func (s *BookingTestSuite) SetupSuite() {
	config.InitLogger()

	log.Info("============================================")
	log.Info("Setting up Booking Test Suite")
	log.Info("============================================")

	caps := selenium.Capabilities{"browserName": "chrome"}

	var err error
	s.wd, err = selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to Selenium Server")
	}

	log.Info("✅ Browser started successfully")
}

func (s *BookingTestSuite) TearDownSuite() {
	log.Info("============================================")
	log.Info("Tearing down Booking Test Suite")
	log.Info("============================================")

	if s.wd != nil {
		s.wd.Quit()
		log.Info("✅ Browser closed successfully")
	}
}

func (s *BookingTestSuite) SetupTest() {
	testName := s.T().Name()

	log.WithField("test", testName).Info("Setting up test")

	if err := s.wd.Get("https://blazedemo.com/"); err != nil {
		log.WithError(err).Error("Failed to navigate to BlazeDemo")
		s.T().Fatal(err)
	}

	log.WithField("test", testName).Info("✅ Navigated to BlazeDemo homepage")
}

func (s *BookingTestSuite) TearDownTest() {
	testName := s.T().Name()

	log.WithField("test", testName).Info("Tearing down test")

	if s.T().Failed() {
		log.WithField("test", testName).Error("❌ Test FAILED - taking screenshot")

		// Сделать скриншот
		screenshot := s.takeScreenshot(fmt.Sprintf("failure_%s", testName))

		// Прикрепить к Allure отчету
		if screenshot != nil {
			allure.AddAttachment("Failure Screenshot", allure.ImagePng, screenshot)
		}
	} else {
		log.WithField("test", testName).Info("✅ Test PASSED")
	}
}

func (s *BookingTestSuite) TestBookingFlight() {
	allure.Test(s.T(),
		allure.Description("Complete flight booking from Paris to London"),
		allure.Epic("E-commerce"),
		allure.Feature("Flight Booking"),
		allure.Story("Book International Flight"),

		allure.Action(func() {
			// ШАГ 1: Проверить title
			log.Info("STEP 1: Verifying page title")

			title, err := s.wd.Title()
			if err != nil {
				log.WithError(err).Error("Failed to get page title")

				screenshot := s.takeScreenshot("title_error")
				if screenshot != nil {
					allure.AddAttachment("Title Error", allure.ImagePng, screenshot)
				}

				s.T().Fatal(err)
			}

			log.WithField("title", title).Info("Page title retrieved")
			assert.Equal(s.T(), "BlazeDemo", title, "Title should be BlazeDemo")

			// Скриншот успешного шага
			screenshot := s.takeScreenshot("homepage")
			if screenshot != nil {
				allure.AddAttachment("1. Homepage", allure.ImagePng, screenshot)
			}

			// ШАГ 2: Выбрать город отправления
			log.Info("STEP 2: Selecting departure city: Paris")

			bookingPage := pages.NewBookingPage(s.wd)

			err = bookingPage.SelectDepartureCity("Paris")
			if err != nil {
				log.WithError(err).Error("Failed to select departure city")

				screenshot := s.takeScreenshot("departure_error")
				if screenshot != nil {
					allure.AddAttachment("Departure Error", allure.ImagePng, screenshot)
				}

				s.T().Fatal(err)
			}

			// Скриншот после выбора
			screenshot = s.takeScreenshot("departure_selected")
			if screenshot != nil {
				allure.AddAttachment("2. Departure Selected - Paris", allure.ImagePng, screenshot)
			}

			// ШАГ 3: Выбрать город назначения
			log.Info("STEP 3: Selecting destination city: London")

			err = bookingPage.SelectDestinationCity("London")
			if err != nil {
				log.WithError(err).Error("Failed to select destination city")

				screenshot := s.takeScreenshot("destination_error")
				if screenshot != nil {
					allure.AddAttachment("Destination Error", allure.ImagePng, screenshot)
				}

				s.T().Fatal(err)
			}

			// Скриншот после выбора
			screenshot = s.takeScreenshot("destination_selected")
			if screenshot != nil {
				allure.AddAttachment("3. Destination Selected - London", allure.ImagePng, screenshot)
			}

			// ШАГ 4: Нажать Find Flights
			log.Info("STEP 4: Clicking Find Flights button")

			err = bookingPage.ClickFindFlights()
			if err != nil {
				log.WithError(err).Error("Failed to click Find Flights")

				screenshot := s.takeScreenshot("find_flights_error")
				if screenshot != nil {
					allure.AddAttachment("Find Flights Error", allure.ImagePng, screenshot)
				}

				s.T().Fatal(err)
			}

			// Финальный скриншот
			screenshot = s.takeScreenshot("flights_found")
			if screenshot != nil {
				allure.AddAttachment("4. Flights Found", allure.ImagePng, screenshot)
			}

			log.Info("✅ Flight booking test completed successfully")
		}),
	)
}

func TestBookingTestSuite(t *testing.T) {
	suite.Run(t, new(BookingTestSuite))
}
