package goiex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Symbol        string
	CompanyName   string
	Exchange      string `json:"primaryExchange"`
	Sector        string
	Open          float64
	Close         float64
	High          float64
	Low           float64
	Volume        int     `json:"latestVolume"`
	Price         float64 `json:"latestPrice"`
	Change        float64
	ChangePercent float64
	AverageVolume int `json:"avgTotalVolume"`
	MarketCap     int
	PeRatio       float64
	YearHigh      float64 `json:"week52High"`
	YearLow       float64 `json:"week52Low"`
	YtdChange     float64
}

func Get(ticker string) (Stock, error) {
	url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/quote", ticker)
	resp, err := http.Get(url)

	if err != nil {
		return Stock{}, err
	}

	s := Stock{}

	err = json.NewDecoder(resp.Body).Decode(&s)

	if err != nil {
		return Stock{}, fmt.Errorf("Unknown ticker %s", ticker)
	}

	return s, nil
}
