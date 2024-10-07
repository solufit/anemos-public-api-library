package libanemos

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
	reported_time     string
}

type Earthquake struct {
	event_id          string
	entry_id          string
	editorial_office  string
	publishing_office string
	category          string
	datetime          string
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
	datetime          string
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
	reported_at    string
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
	reported_at    string
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
	reported_at    string
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

func (m WeatherWarninglist) WeatherWarningFilter(filterOption FilterOptions) WeatherWarninglist {
	filteredData := make([]WeatherWarning, 0)
	for _, warning := range m.data {
		// Apply filter conditions here
		if warning.reported_at >= filterOption.StartTime.String() && warning.reported_at <= filterOption.EndTime.String() {
			filteredData = append(filteredData, warning)
		}
	}
	return WeatherWarninglist{data: filteredData}
}

func (m WeatherForecastlist) WeatherForecastFilter(filterOption FilterOptions) WeatherForecastlist {
	filteredData := make([]WeatherForecast, 0)
	for _, forecast := range m.data {
		if forecast.reported_at >= filterOption.StartTime.String() && forecast.reported_at <= filterOption.EndTime.String() {
			filteredData = append(filteredData, forecast)
		}
	}
	return WeatherForecastlist{data: filteredData}
}

func (m WeatherEarthquakelist) WeatherEarthquakeFilter(filterOption FilterOptions) WeatherEarthquakelist {
	filteredData := make([]WeatherEarthquake, 0)
	for _, earthquake := range m.data {
		// Apply filter conditions here
		if earthquake.reported_at >= filterOption.StartTime.String() && earthquake.reported_at <= filterOption.EndTime.String() {
			filteredData = append(filteredData, earthquake)
		}
	}
	return WeatherEarthquakelist{data: filteredData}
}
