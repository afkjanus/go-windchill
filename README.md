# go-windchill

This package calculates the current wind chill temperature based on the air temperature, and the wind speed.
Using the function WindChill.

The wind chill can only be calculated for air temperatures below 10 °C and wind speeds above 4.8 km/h.
If an air temperature or wind speed is provided that is outside of these bounds an error will be returned.
If an error is returned the returned wind chill temperature will be the air temperature.
Therefore, a temperature is provided for all input values.

The calculations are based on: [https://en.wikipedia.org/wiki/Wind_chill](https://en.wikipedia.org/wiki/Wind_chill])