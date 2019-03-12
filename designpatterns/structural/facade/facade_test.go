package facade

import (
	"bytes"
	"io"
	"testing"
)

func getMockData() io.Reader {
	response := `{
		"id": 3117735,
		"name": "Madrid",
		"coord": {
			"lon": -3.7,
			"lat": 40.42
		},
		"sys": {
			"country": "ES"
		}
	}`

	return bytes.NewReader([]byte(response))
}

func TestOpenWeatherMap_responseParser(t *testing.T) {
	r := getMockData()
	openWeatherMap := CurrentWeatherData{APIKey: ""}
	weather, err := openWeatherMap.responseParser(r)
	if err != nil {
		t.Fatal(err)
	}
	if weather.ID != 3117735 {
		t.Errorf("Madrid id is 3117735, not %d\n", weather.ID)
	}
}

func TestOpenWeatherMap_GetByGeoCoordinates(t *testing.T) {
	openWeatherMap := CurrentWeatherData{APIKey: "f352019b83038c32f99f679ad97af36a"}
	weather, err := openWeatherMap.GetByCityAndCountryCode("Madrid", "ES")
	if err != nil {
		t.Fatal(err)
	}

	if !(weather.Main.Temp > 0) {
		t.Error("Data was not retrieved")
	}

}
