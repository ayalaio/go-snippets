package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, cc string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type CurrentWeatherData struct {
	APIKey string
}

type Weather struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Main struct {
		Temp float32 `json:"temp"`
	} `json: "main"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Sys struct {
		Country string `json:"country"`
	}
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*Weather, error) {
	return c.doRequest(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIKey),
	)
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(ct, cc string) (*Weather, error) {
	return c.doRequest(
		fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s,%s&APPID=%s", ct, cc, c.APIKey),
	)
}

func (c *CurrentWeatherData) doRequest(uri string) (*Weather, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		byt, errMsg := ioutil.ReadAll(resp.Body)
		if errMsg == nil {
			errMsg = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code %d\n Error: %s", resp.StatusCode, errMsg)
	}

	weather, err := c.responseParser(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return weather, nil

}

func (c *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := &Weather{}
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
