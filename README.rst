.. image:: https://travis-ci.com/AndrewRPorter/goiex.svg?branch=master
    :target: https://travis-ci.com/AndrewRPorter/goiex
    
=====
goiex
=====

**Note: the IEX** `api <https://iextrading.com/developers/docs/>`_ **now requires a key for authentication, thus breaking this project.**

A go interface for accessing IEX finanical information. Data provided for free 
by `IEX <https://iextrading.com/developer/>`_. View `IEX’s Terms of Use 
<https://iextrading.com/api-exhibit-a/>`_.


Installation
------------

.. code::
   
      go get github.com/AndrewRPorter/goiex

Usage
-----

.. code:: go

   package main
   
   import (
      "fmt"
      "github.com/AndrewRPorter/goiex"
   )

   func main() {   
      s, err := goiex.Get("AAPL")
      
      if err != nil {
         fmt.Errorf("Unable to fetch ticker: %s", "AAPL")
      }
      
      fmt.Println(s.Price)
      
      // fetch company information
      c, err := s.GetCompany()
      
      if err != nil {
         fmt.Errorf("Unable to fetch company data for: %s", s.Symbol)
      }
      fmt.Println(c.Description)
   }
   
**Available Methods**

- ``Get()``
- ``GetCompany()``

**Available Fields**

- ``Symbol``
- ``CompanyName``
- ``Exchange``
- ``Sector``
- ``Open``
- ``Close``
- ``High``
- ``Low``
- ``Volume``
- ``Price``
- ``Change``
- ``ChangePercent``
- ``AverageVolume``
- ``MarketCap``
- ``PeRatio``
- ``YearHigh``
- ``YearLow``
- ``YtdChange``

- ``Symbol``
- ``CompanyName``
- ``Exchange``
- ``Industry``
- ``Website``
- ``Description``
- ``CEO``
- ``Sector``
