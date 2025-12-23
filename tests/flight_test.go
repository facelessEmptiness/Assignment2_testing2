package tests

import (
	"testing"

	"github.com/facelessEmptiness/Assignment2_testing/pages"
	"github.com/facelessEmptiness/Assignment2_testing/utils"
	"github.com/stretchr/testify/assert"
)

func TestFlightBooking(t *testing.T) {
	wd, service, err := utils.SetupWebDriver("chrome")
	if err != nil {
		t.Fatalf("Failed to setup WebDriver: %v", err)
	}
	defer wd.Quit()
	defer service.Stop()

	flightPage := &pages.FlightPage{BasePage: pages.BasePage{WebDriver: wd}}

	// Используем демо сайт для бронирования рейсов
	t.Run("TestFlightBookingWithTitleCheckpoint", func(t *testing.T) {
		err := flightPage.NavigateTo("https://blazedemo.com")
		assert.NoError(t, err)

		// Проверка заголовка (checkpoint)
		title, err := flightPage.GetPageTitle()
		assert.NoError(t, err)
		assert.Contains(t, title, "BlazeDemo", "Page title should contain BlazeDemo")

		// Выбор пунктов отправления и назначения
		err = flightPage.SelectDepartureCity("Paris")
		assert.NoError(t, err)

		err = flightPage.SelectDestinationCity("London")
		assert.NoError(t, err)

		// Поиск рейсов
		err = flightPage.FindFlights()
		assert.NoError(t, err)

		// Проверка наличия рейсов
		hasFlights, err := flightPage.HasAvailableFlights()
		assert.NoError(t, err)
		assert.True(t, hasFlights, "Should have available flights")

		// Выбор первого доступного рейса
		err = flightPage.ChooseFirstFlight()
		assert.NoError(t, err)

		// Заполнение формы бронирования
		err = flightPage.FillPassengerInfo(
			"John",
			"Doe",
			"john.doe@example.com",
			"1234567890",
			"123 Main St",
		)
		assert.NoError(t, err)

		// Проверка перед покупкой
		totalPrice, err := flightPage.GetTotalPrice()
		assert.NoError(t, err)
		assert.NotEmpty(t, totalPrice, "Should display total price")

		// Завершение бронирования
		err = flightPage.PurchaseFlight()
		assert.NoError(t, err)

		// Проверка подтверждения
		isConfirmed, err := flightPage.IsBookingConfirmed()
		assert.NoError(t, err)
		assert.True(t, isConfirmed, "Booking should be confirmed")

		// Сохраняем скриншот подтверждения
		utils.TakeScreenshot(wd, "flight_booking_confirmation.png")

		// Проверка ID бронирования
		bookingID, err := flightPage.GetBookingID()
		assert.NoError(t, err)
		assert.NotEmpty(t, bookingID, "Should have booking ID")

		t.Logf("Flight booked successfully! Booking ID: %s", bookingID)
	})
}
