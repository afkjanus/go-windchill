package windchill

import (
	"errors"
	"math"
)

// WindChill returns the current wind chill temperature based on the air temperature, and the wind speed.
// The wind chill can only be calculated for air temperatures below 10 °C and wind speeds above 4.8 km/h.
// If an air temperature or wind speed is provided that is outside of these bounds an error will be returned.
// If an error is returned the returned wind chill temperature will be the air temperature.
// Therefore, a temperature is provided for all input values.
//
// The calculations are based on: https://en.wikipedia.org/wiki/Wind_chill
func WindChill(airTemperature float64, windSpeed float64) (windChillTemperature float64, err error) {
	if airTemperature > 10 {
		return airTemperature, errors.New("the air temperature must be at or above 10 °C to calculate the wind chill temperature")
	}
	if windSpeed <= 4.8 {
		return airTemperature, errors.New("the wind speed must be above 4.8 km/h to calculate teh wind chill temperature")
	}

	return 13.12 + 0.6215*airTemperature + (0.3965*airTemperature-11.37)*math.Pow(windSpeed, 0.16), nil
}
