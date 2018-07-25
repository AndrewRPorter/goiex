package goiex

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL string = "https://api.iextrading.com/1.0/stock"

// Stock provides an interface for accessing ticker data parsed from IEX
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

// Company provides and interface for accessing company data parsed from IEX
type Company struct {
	Symbol      string
	CompanyName string
	Exchange    string
	Industry    string
	Website     string
	Description string
	CEO         string
	Sector      string
}

// GetCompany interpolates Company struct with JSON requests data from IEX
func (s Stock) GetCompany() (Company, error) {
	url := fmt.Sprintf("%s/%s/company", baseURL, s.Symbol)
	resp, err := http.Get(url)

	if err != nil {
		return Company{}, err
	}

	c := Company{}

	err = json.NewDecoder(resp.Body).Decode(&c)

	if err != nil {
		return Company{}, fmt.Errorf("Unknown ticker %s", s.Symbol)
	}

	return c, nil
}

// Get interpolates Stock struct with JSON request data from IEX
func Get(ticker string) (Stock, error) {
	url := fmt.Sprintf("%s/%s/quote", baseURL, ticker)
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
