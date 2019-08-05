package ratara

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

var rng *rand.Rand

type TokyoCities struct {
	Tokyo []struct {
		Citycode string `json:"citycode"`
		City     string `json:"city"`
	} `json:"tokyo"`
}

type Countries struct {
	Country []struct {
		Name string `json:"name"`
	} `json:"country"`
}

func initRand() {
	if rng != nil {
		return
	}
	rng = rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
}

func tokyo() *TokyoCities {
	tokyo := new(TokyoCities)
	cities, err := Assets.Open("/data/Tokyo.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(cities).Decode(&tokyo); err != nil {
		panic(err)
	}
	return tokyo
}

func country() *Countries {
	country := new(Countries)
	countries, err := Assets.Open("/data/RestaurantCountry.json")
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(countries).Decode(&country); err != nil {
		panic(err)
	}

	return country
}

// Place return the place where you eat at.
func Place() string {
	initRand()
	tokyo := tokyo()
	cityNum := len(tokyo.Tokyo)
	num := rng.Intn(cityNum)
	return tokyo.Tokyo[num].City
}

// FoodCountry return the country which you shoud eat which country food you eat
func FoodCountry() string {
	initRand()
	country := country()
	countryNum := len(country.Country)
	num := rng.Intn(countryNum)
	return country.Country[num].Name
}
