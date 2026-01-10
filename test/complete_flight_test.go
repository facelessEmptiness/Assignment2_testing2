package test

import (
	"testing"

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
		s.T().Fatal(err) // ← Останавливает ТОЛЬКО этот тест
	}

	log.WithField("test", testName).Info("✅ Navigated to BlazeDemo homepage")
}

func (s *BookingTestSuite) TearDownTest() {
	testName := s.T().Name()

	log.WithField("test", testName).Info("Tearing down test")

	if s.T().Failed() {
		log.WithField("test", testName).Error("❌ Test FAILED")

	} else {
		log.WithField("test", testName).Info("✅ Test PASSED")
	}
}

func (s *BookingTestSuite) TestBookingFlight() {
	log.Info("Starting flight booking test")

	// Проверка title
	title, err := s.wd.Title()
	if err != nil {
		log.WithError(err).Error("Failed to get page title")
		s.T().Fatal(err)
	}

	log.WithField("title", title).Info("Page title retrieved")
	assert.Equal(s.T(), "BlazeDemo", title, "Title should be BlazeDemo")

	// Создать Page Object
	bookingPage := pages.NewBookingPage(s.wd)

	// Выбрать Paris
	err = bookingPage.SelectDepartureCity("Paris")
	if err != nil {
		log.WithError(err).Error("Failed to select departure city")
		s.T().Fatal(err)
	}

	// Выбрать London
	err = bookingPage.SelectDestinationCity("London")
	if err != nil {
		log.WithError(err).Error("Failed to select destination city")
		s.T().Fatal(err)
	}

	// Нажать Find Flights
	err = bookingPage.ClickFindFlights()
	if err != nil {
		log.WithError(err).Error("Failed to click Find Flights")
		s.T().Fatal(err)
	}

	log.Info("✅ Flight booking test completed successfully")
}

func TestBookingTestSuite(t *testing.T) {
	suite.Run(t, new(BookingTestSuite))
}
