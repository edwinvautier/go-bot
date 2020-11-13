package meteo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

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

// getLocation will get the location details for where this
// application has been run from.
func getLocation() (*Data, error) {
	response, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	r := &Data{}
	if err = json.NewDecoder(response.Body).Decode(&r); err != nil {
		return nil, err
	}
	fmt.Println(r)
	return r, nil
}

// getCurrent gets the current weather for the provided location in
// the units provided.
func getCurrent(l, u, lang string) (*owm.CurrentWeatherData, error) {
	var apiKey = GetKey()

	w, err := owm.NewCurrent(u, lang, apiKey) // Create the instance with the given unit Celsius (metric) with France output
	if err != nil {
		return nil, err
	}
	w.CurrentByName(l) // Get the actual data for the given location
	fmt.Println(w)
	return w, nil
}

// hereHandler will take are of requests coming in for the "/here" route.
func hereHandler(w http.ResponseWriter, r *http.Request) {
	location, err := getLocation()
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	wd, err := getCurrent(location.City, "C", "fr")
	if err != nil {
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}

	fmt.Println(wd)
}

// FindWheatherByCity take the return of the dispatcher and request the requested location.
func FindWheatherByCity(wp *WeatherParams) *WeatherData {
	var apiKey = GetKey()

	log.Info("Recherche du temps")
	w, err := owm.NewCurrent("C", "fr", apiKey) // Celsius (metric) with France output
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	err = w.CurrentByName(wp.Location)

	weatherData := WeatherData{Wheater: &w.Weather}

	return &weatherData
}

// WeatherParams result
type WeatherParams struct {
	Location string
}

type WeatherData struct {
	Wheater *[]owm.Weather
	Main *owm.Weather
}

func (wd *WeatherData) String() string {

	return fmt.Sprintf("weather:%s", wd.Wheater)
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
