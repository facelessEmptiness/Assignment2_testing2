package pages

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tebeka/selenium"
)

// BookingPage - Page Object для главной страницы BlazeDemo
type BookingPage struct {
	wd                    selenium.WebDriver
	selectDepartureCity   string
	selectDestinationCity string
	findFlightBtn         string
}

// NewBookingPage - конструктор Page Object
func NewBookingPage(wd selenium.WebDriver) *BookingPage {
	return &BookingPage{
		wd:                    wd,
		selectDepartureCity:   "select[name='fromPort']",
		selectDestinationCity: "select[name='toPort']",
		findFlightBtn:         "input[type='submit']",
	}
}

// SelectDepartureCity - выбрать город отправления
func (b *BookingPage) SelectDepartureCity(city string) error {
	log.WithFields(log.Fields{
		"page":   "BookingPage",
		"action": "SelectDepartureCity",
		"city":   city,
	}).Info("Selecting departure city")

	// Найти dropdown
	departureSelect, err := b.wd.FindElement(selenium.ByCSSSelector, b.selectDepartureCity)
	if err != nil {
		log.WithFields(log.Fields{
			"selector": b.selectDepartureCity,
			"error":    err.Error(),
		}).Error("Failed to find departure dropdown")
		return fmt.Errorf("failed to find departure dropdown: %v", err)
	}

	// Кликнуть dropdown
	if err := departureSelect.Click(); err != nil {
		log.WithError(err).Error("Failed to click departure dropdown")
		return fmt.Errorf("failed to click departure dropdown: %v", err)
	}

	// Найти опцию с городом
	xpath := fmt.Sprintf("//select[@name='fromPort']//option[@value='%s']", city)
	log.WithField("xpath", xpath).Debug("Searching for city option")

	option, err := b.wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		log.WithFields(log.Fields{
			"xpath": xpath,
			"city":  city,
		}).Error("Failed to find city option")
		return fmt.Errorf("failed to find city %s: %v", city, err)
	}

	// Кликнуть опцию
	if err := option.Click(); err != nil {
		log.WithError(err).Error("Failed to click city option")
		return fmt.Errorf("failed to select city: %v", err)
	}

	log.WithField("city", city).Info("✅ Departure city selected successfully")
	return nil
}

// SelectDestinationCity - выбрать город назначения
func (b *BookingPage) SelectDestinationCity(city string) error {
	log.WithFields(log.Fields{
		"page":   "BookingPage",
		"action": "SelectDestinationCity",
		"city":   city,
	}).Info("Selecting destination city")

	// Найти dropdown
	destinationSelect, err := b.wd.FindElement(selenium.ByCSSSelector, b.selectDestinationCity)
	if err != nil {
		log.WithFields(log.Fields{
			"selector": b.selectDestinationCity,
			"error":    err.Error(),
		}).Error("Failed to find destination dropdown")
		return fmt.Errorf("failed to find destination dropdown: %v", err)
	}

	// Кликнуть dropdown
	if err := destinationSelect.Click(); err != nil {
		log.WithError(err).Error("Failed to click destination dropdown")
		return fmt.Errorf("failed to click destination dropdown: %v", err)
	}

	// Найти опцию с городом
	xpath := fmt.Sprintf("//select[@name='toPort']//option[@value='%s']", city)
	log.WithField("xpath", xpath).Debug("Searching for destination city option")

	option, err := b.wd.FindElement(selenium.ByXPATH, xpath)
	if err != nil {
		log.WithFields(log.Fields{
			"xpath": xpath,
			"city":  city,
		}).Error("Failed to find destination city option")
		return fmt.Errorf("failed to find destination city %s: %v", city, err)
	}

	// Кликнуть опцию
	if err := option.Click(); err != nil {
		log.WithError(err).Error("Failed to click destination city option")
		return fmt.Errorf("failed to select destination city: %v", err)
	}

	log.WithField("city", city).Info("✅ Destination city selected successfully")
	return nil
}

// ClickFindFlights - нажать кнопку поиска рейсов
func (b *BookingPage) ClickFindFlights() error {
	log.WithField("page", "BookingPage").Info("Clicking Find Flights button")

	// Найти кнопку
	findFlightsBtn, err := b.wd.FindElement(selenium.ByCSSSelector, b.findFlightBtn)
	if err != nil {
		log.WithFields(log.Fields{
			"selector": b.findFlightBtn,
			"error":    err.Error(),
		}).Error("Failed to find Find Flights button")
		return fmt.Errorf("failed to find Find Flights button: %v", err)
	}

	// Кликнуть кнопку
	if err := findFlightsBtn.Click(); err != nil {
		log.WithError(err).Error("Failed to click Find Flights button")
		return fmt.Errorf("failed to click Find Flights button: %v", err)
	}

	log.Info("✅ Find Flights button clicked successfully")

	// НЕТ time.Sleep! Если нужно ждать - добавим явное ожидание позже
	return nil
}
