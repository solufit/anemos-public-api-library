package libanemos

import "time"

type Forecast struct {
	weather_today     string
	weather_tommorow  string
	max_temp          int
	min_temp          int
	rain_percent_now  int
	rain_percent_6h   int
	rain_percent_12h  int
	rain_percent_18h  int
	rain_percent_24h  int
	publishing_office string
	reported_time     time.Time
}

type Earthquake struct {
	event_id          string
	entry_id          string
	editorial_office  string
	publishing_office string
	category          string
	datetime          time.Time
	headline          string
	hypocenter        string
	region_code       string
	max_int           string
	magnitude         string
}

type Warning struct {
	entryid           string
	editorial_office  string
	publishing_office string
	category          string
	datetime          time.Time
	headline          string
	pref              string
}

type WeatherForecast struct {
	id             string
	object_type    string
	areacode       string
	title          string
	status         string
	detail         Forecast
	reported_at    time.Time
	info_domain    string
	info_object_id string
}

type WeatherEarthquake struct {
	id             string
	object_type    string
	areacode       string
	title          string
	status         string
	detail         Earthquake
	reported_at    time.Time
	info_domain    string
	info_object_id string
}

type WeatherWarning struct {
	id             string
	object_type    string
	areacode       string
	title          string
	status         string
	detail         Warning
	reported_at    time.Time
	info_domain    string
	info_object_id string
}

type WeatherWarninglist struct {
	data []WeatherWarning
}

type WeatherEarthquakelist struct {
	data []WeatherEarthquake
}

type WeatherForecastlist struct {
	data []WeatherForecast
}

func (m WeatherWarninglist) Filter(filterOption FilterOptions) WeatherWarninglist {
	filteredData := make([]WeatherWarning, 0)
	filteredData = append(filteredData, m.data...)

	if filterOption.StartTime != nil {
		tmp_data := make([]WeatherWarning, 0)
		for _, weather := range m.data {
			if filterOption.StartTime.IsSome() {
				startTime, _ := filterOption.StartTime.Take()
				if startTime.Before(weather.reported_at) {
					tmp_data = append(tmp_data, weather)
				}
			}
		}
		filteredData = tmp_data
	}

	if filterOption.EndTime != nil {
		tmp_data := make([]WeatherWarning, 0)
		for _, weather := range m.data {
			if filterOption.EndTime.IsSome() {
				EndTime, _ := filterOption.EndTime.Take()
				if EndTime.After(weather.reported_at) {
					tmp_data = append(tmp_data, weather)
				}
			}
		}
		filteredData = tmp_data
	}
	return WeatherWarninglist{data: filteredData}
}

func (m WeatherForecastlist) Filter(filterOption FilterOptions) WeatherForecastlist {
	filteredData := make([]WeatherForecast, 0)

	filteredData = append(filteredData, m.data...)

	if filterOption.StartTime != nil {
		tmp_data := make([]WeatherForecast, 0)
		for _, weather := range m.data {
			if filterOption.StartTime.IsSome() {
				startTime, _ := filterOption.StartTime.Take()
				if startTime.Before(weather.reported_at) {
					tmp_data = append(tmp_data, weather)
				}
			}
		}
		filteredData = tmp_data
	}

	if filterOption.EndTime != nil {
		tmp_data := make([]WeatherForecast, 0)
		for _, weather := range m.data {
			if filterOption.EndTime.IsSome() {
				EndTime, _ := filterOption.EndTime.Take()
				if EndTime.After(weather.reported_at) {
					tmp_data = append(tmp_data, weather)
				}
			}
		}
		filteredData = tmp_data
	}
	return WeatherForecastlist{data: filteredData}
}

func (m WeatherEarthquakelist) Filter(filterOption FilterOptions) WeatherEarthquakelist {
	filteredData := make([]WeatherEarthquake, 0)
	filteredData = append(filteredData, m.data...)

	if filterOption.StartTime != nil {
		tmp_data := make([]WeatherEarthquake, 0)
		for _, earthquake := range m.data {
			if filterOption.StartTime.IsSome() {
				startTime, _ := filterOption.StartTime.Take()
				if startTime.Before(earthquake.reported_at) {
					tmp_data = append(tmp_data, earthquake)
				}
			}
		}
		filteredData = tmp_data
	}

	if filterOption.EndTime != nil {
		tmp_data := make([]WeatherEarthquake, 0)
		for _, earthquake := range m.data {
			if filterOption.EndTime.IsSome() {
				EndTime, _ := filterOption.EndTime.Take()
				if EndTime.After(earthquake.reported_at) {
					tmp_data = append(tmp_data, earthquake)
				}
			}
		}
		filteredData = tmp_data
	}

	return WeatherEarthquakelist{data: filteredData}
}
