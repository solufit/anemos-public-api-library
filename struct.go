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

type WeatherInfoFilter struct {
	Number    int    `json:"number"`
	PostCode  string `json:"post_code"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Filter    string `json:"filter" enum:"weather-forecast,weather-warning,earthquake"`
}
