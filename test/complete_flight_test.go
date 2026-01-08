package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/tebeka/selenium"
)

type BookingTestSuite struct {
	suite.Suite
	wd selenium.WebDriver
}

func (s *BookingTestSuite) SetupSuite() {
	fmt.Println("Setting up Suite")
	caps := selenium.Capabilities{"browserName": "chrome"}

	var err error
	s.wd, err = selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	fmt.Println("Browser Opened Successfully")
}
func (s *BookingTestSuite) TearDownSuite() {
	fmt.Println("Tearing Down Suite")
	s.wd.Quit()
	fmt.Println("Browser Closed Successfully")
}
func (s *BookingTestSuite) SetupTest() {
	testName := s.T().Name()
	fmt.Printf("Setting up Test", testName)
	if err := s.wd.Get("https://blazedemo.com/"); err != nil {
		s.T().Fatal(err)
	}
	fmt.Println("✓ Navigated to BlazeDemo homepage\n")
}
func (s *BookingTestSuite) TearDownTest() {
	fmt.Println("Tearing Down Test")
	if s.T().Failed() {
		fmt.Println("Test failed!")
	} else {
		fmt.Println("Test passed!")
	}

}
func (s *BookingTestSuite) TestBookingFlight() {
	fmt.Println("Starting Booking Flight")
	title, err := s.wd.Title()
	if err != nil {
		s.T().Fatal(err)
	}
	assert.Equal(s.T(), "BlazeDemo", title, "Title should be BlazeDemo")
	fmt.Println("Test completed")
}
func TestBookinTestSuite(t *testing.T) {
	suite.Run(t, new(BookingTestSuite))
}
