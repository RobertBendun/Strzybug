package weather

import "time"

type Response struct {
	Latitude float32 `json:"lat"`
	Longitude float32 `json:"lon"`
	Timezone string
	TimezoneOffset uint
	Daily []Day
}

type Day struct {
	Time int64 `json:"dt"`
	Sunrise int64
	Sunset int64
	Moonrise int64
	Moonset int64
	MoonPhase float32 `json:"moon_phase"`
	Temp Temp
	Feels Temp
	Pressure float32

	// Humidity in %
	Humidity float32

	// Speed of wind in m/s
	WindSpeed float32

	// Wind direction in degrees
	WindDegrees float32

	// Wind gust in m/s
	WindGust float32

	// Cloudiness in %
	Clouds float32

	// Probability of precipitation, from 0 to 1
	Precipitation float32 `json:"pop"`

	Weather []Weather
}

type Temp struct {
	Day float32
	Night float32
	Morning float32 `json:"morn"`
	Evening float32 `json:"eve"`
}

type Weather struct {
	Id          uint

	// Group of weather, like Rain, Snow, Extreme etc.
	Main        string

	// Localized description of weather
	Description string

	Icon        string
}

func (r Response) FindFirstDate() time.Time {
	return time.Unix(r.Daily[0].Time, 0)
}
