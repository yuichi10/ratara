package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

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

func main() {
	tokyo := tokyo()
	country := country()
	cityNum := len(tokyo.Tokyo)
	countryNum := len(country.Country)
	rng := rand.New(mt19937.New())
	rng.Seed(time.Now().UnixNano())
	cil := rng.Intn(cityNum)
	col := rng.Intn(countryNum)
	fmt.Println(tokyo.Tokyo[cil].City)
	fmt.Println(country.Country[col].Name)
}
