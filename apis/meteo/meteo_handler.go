package meteo

import (
	"strings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	owm "github.com/briandowns/openweathermap"
)

// GetKey get OWN_API_KEY
func GetKey() string {
	apiKey, tokenExist := os.LookupEnv("OWN_API_KEY")
	if !tokenExist {
		log.Error("Missing environment variable OWN_API_KEY")
		return ""
	}
	return apiKey
}

// URL is a constant that contains where to find the IP locale info
const URL = "http://ip-api.com/json"

// GetLocation will get the location details for where this application has been run from.
func GetLocation() (*Data, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	loc := &Data{}
	if err = json.Unmarshal(result, &loc); err != nil {
		return nil, err
	}
	return loc, nil
}

// getCurrent gets the current weather for the provided location in the units provided.
func getCurrent(l, u, lang string) (*owm.CurrentWeatherData, error) {
	var apiKey = GetKey()

	w, err := owm.NewCurrent(u, lang, apiKey) // Create the instance with the given unit Celsius (metric) with France output
	if err != nil {
		return nil, err
	}
	w.CurrentByName(l) // Get the actual data for the given location
	return w, nil
}

// GetHereHandler requested by IP API.
func GetHereHandler() (*WeatherData, error) {
	location, err := GetLocation()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	wd, err := getCurrent(location.City, "C", "fr")
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	weatherData := WeatherData{Main: &wd.Main}

	return &weatherData, err
}

// FindWheatherByCity requested by city.
func FindWheatherByCity(wp *WeatherParams) *WeatherData {
	var apiKey = GetKey()

	w, err := owm.NewCurrent("C", "fr", apiKey) // Celsius (metric) with France output
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	
	err = w.CurrentByName(strings.ToLower(wp.Location))

	weatherData := WeatherData{Main: &w.Main}

	return &weatherData
}

// WeatherParams result
type WeatherParams struct {
	Location string
}

// WeatherData result
type WeatherData struct {
	Main *owm.Main
}

func (wd *WeatherData) String() string {
	return fmt.Sprintf("Température : %v°C \n Ressenti: %v°C", wd.Main.Temp, wd.Main.FeelsLike)
}

// Data will hold the result of the query to get the IP
// address of the caller.
type Data struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	ISP         string  `json:"isp"`
	ORG         string  `json:"org"`
	AS          string  `json:"as"`
	Message     string  `json:"message"`
	Query       string  `json:"query"`
}
