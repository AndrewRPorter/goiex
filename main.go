package main

import (
   "fmt"
   "log"
   "encoding/json"
   "github.com/levigross/grequests"
)

type Stock struct{
   Symbol string
   CompanyName string
   PrimaryExchange string
   Sector string
   CalculationPrice string
   Open float64
   OpenTime int
   Close float64
   CloseTime int
   High float64
   Low float64
   LatestPrice float64
   LatestSource string
   LatestTime string
   LatestUpdate int
   LatestVolume int
   IexRealtimePrice float64
   IexRealtimeSize int
   IexLastUpdated int
   DelayedPrice float64
   DelayedPriceTime int
   ExtendedPrice float64
   ExtendedChange float64
   ExtendedChangePercent float64
   ExtendedPriceTime int
   Change float64
   ChangePercent float64
   IexMarketPercent float64
   IexVolume int
   AvgTotalVolume int
   IexBidPrice float64
   IexBidSize int
   IexAskPrice float64
   IexAskSize int
   MarketCap int
   PeRatio float64
   Week52High float64
   Week52Low float64
   YtdChange float64
}

func fetch(ticker string) Stock {
   url := fmt.Sprintf("https://api.iextrading.com/1.0/stock/%s/quote", ticker)
   resp, err := grequests.Get(url, nil)

   if err != nil {
    log.Fatalln("Unable to make request: ", err)
   }

   s := Stock{}

   err = json.NewDecoder(resp).Decode(&s)

   if err != nil {
      log.Fatalln("Unable to decode JSON: ", err)
   }

   return s
}

func (s *Stock) get_price() float64 {
   return s.LatestPrice
}

func (s *Stock) get_symbol() string {
   return s.Symbol
}

func main() {
   stock := fetch("GEVO")

   fmt.Println(stock.get_symbol())
   fmt.Println(stock.get_price())
}
