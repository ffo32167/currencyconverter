package cbr

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
	"golang.org/x/net/html/charset"
)

type Cbr struct {
	connStr    string
	currencies string
	client     http.Client
}

func New(connStr, currencies string, ctxTimeout time.Duration) Cbr {
	return Cbr{
		connStr:    connStr,
		currencies: currencies,
		client:     http.Client{Timeout: ctxTimeout * time.Second},
	}
}

//const str string = `<?xml version="1.0" encoding="windows-1251"?><ValCurs Date="02.03.2002" name="Foreign Currency Market"><Valute ID="R01010"><NumCode>036</NumCode><CharCode>AUD</CharCode><Nominal>1</Nominal><Name>Австралийский доллар</Name><Value>16,0102</Value></Valute><Valute ID="R01035"><NumCode>826</NumCode><CharCode>GBP</CharCode><Nominal>1</Nominal><Name>Фунт стерлингов Соединенного королевства</Name><Value>43,8254</Value></Valute><Valute ID="R01090"><NumCode>974</NumCode><CharCode>BYR</CharCode><Nominal>1000</Nominal><Name>Белорусских рублей</Name><Value>18,4290</Value></Valute><Valute ID="R01215"><NumCode>208</NumCode><CharCode>DKK</CharCode><Nominal>10</Nominal><Name>Датских крон</Name><Value>36,1010</Value></Valute><Valute ID="R01235"><NumCode>840</NumCode><CharCode>USD</CharCode><Nominal>1</Nominal><Name>Доллар США</Name><Value>30,9436</Value></Valute><Valute ID="R01239"><NumCode>978</NumCode><CharCode>EUR</CharCode><Nominal>1</Nominal><Name>Евро</Name><Value>26,8343</Value></Valute><Valute ID="R01310"><NumCode>352</NumCode><CharCode>ISK</CharCode><Nominal>100</Nominal><Name>Исландских крон</Name><Value>30,7958</Value></Valute><Valute ID="R01335"><NumCode>398</NumCode><CharCode>KZT</CharCode><Nominal>100</Nominal><Name>Казахстанских тенге</Name><Value>20,3393</Value></Valute><Valute ID="R01350"><NumCode>124</NumCode><CharCode>CAD</CharCode><Nominal>1</Nominal><Name>Канадский доллар</Name><Value>19,3240</Value></Valute><Valute ID="R01535"><NumCode>578</NumCode><CharCode>NOK</CharCode><Nominal>10</Nominal><Name>Норвежских крон</Name><Value>34,7853</Value></Valute><Valute ID="R01589"><NumCode>960</NumCode><CharCode>XDR</CharCode><Nominal>1</Nominal><Name>СДР (специальные права заимствования)</Name><Value>38,4205</Value></Valute><Valute ID="R01625"><NumCode>702</NumCode><CharCode>SGD</CharCode><Nominal>1</Nominal><Name>Сингапурский доллар</Name><Value>16,8878</Value></Valute><Valute ID="R01700"><NumCode>792</NumCode><CharCode>TRL</CharCode><Nominal>1000000</Nominal><Name>Турецких лир</Name><Value>22,2616</Value></Valute><Valute ID="R01720"><NumCode>980</NumCode><CharCode>UAH</CharCode><Nominal>10</Nominal><Name>Украинских гривен</Name><Value>58,1090</Value></Valute><Valute ID="R01770"><NumCode>752</NumCode><CharCode>SEK</CharCode><Nominal>10</Nominal><Name>Шведских крон</Name><Value>29,5924</Value></Valute><Valute ID="R01775"><NumCode>756</NumCode><CharCode>CHF</CharCode><Nominal>1</Nominal><Name>Швейцарский франк</Name><Value>18,1861</Value></Valute><Valute ID="R01820"><NumCode>392</NumCode><CharCode>JPY</CharCode><Nominal>100</Nominal><Name>Японских иен</Name><Value>23,1527</Value></Valute></ValCurs>`

type CbrResponse struct {
	Currency string `xml:"CharCode"`
	Rate     string `xml:"Value"`
	Nominal  string `xml:"Nominal"`
}

func (c Cbr) Rates() ([]internal.Rate, error) {
	resp, err := c.client.Get(c.connStr)
	if err != nil {
		return nil, fmt.Errorf("cant connect with cbr: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got wrong response status from cbr")
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	rates := make([]CbrResponse, 0)
	var date string
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if token == nil {
			return nil, errors.New("cant parse xml")
		}

		switch tp := token.(type) {
		case xml.StartElement:
			if tp.Name.Local == "Valute" {
				var curr CbrResponse
				decoder.DecodeElement(&curr, &tp)
				if strings.Contains(c.currencies, curr.Currency) {
					rates = append(rates, curr)
				}
			} else if tp.Name.Local == "ValCurs" {
				date = tp.Attr[0].Value
			}
		}
	}
	if len(rates) == 0 {
		return nil, errors.New("cant get rates from cbr")
	}
	return toDomain(date, rates)
}

func toDomain(date string, cbr []CbrResponse) ([]internal.Rate, error) {
	result := make([]internal.Rate, len(cbr)+1, len(cbr)+1)

	rateDate, err := time.Parse("02.01.2006", date)
	if err != nil {
		return nil, fmt.Errorf("cant parse date from cbr: %w", err)
	}

	result[0].RateDate = rateDate
	result[0].CurrCode = "RUB"
	result[0].Rate = 1

	for i, v := range cbr {
		dotSeparatedString := strings.Replace(v.Rate, ",", ".", 1)

		rate, err := strconv.ParseFloat(dotSeparatedString, 64)
		if err != nil {
			return nil, errors.New("cant parse rate")
		}
		nominal, err := strconv.ParseFloat(v.Nominal, 64)
		if err != nil {
			return nil, errors.New("cant parse nominal")
		}
		if nominal == 0 {
			return nil, errors.New("nominal cant be 0")
		}
		result[i+1].RateDate = rateDate
		result[i+1].CurrCode = v.Currency
		result[i+1].Rate = math.Round(rate/nominal*1000000) / 1000000
	}
	return result, nil
}
