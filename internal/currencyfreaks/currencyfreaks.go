package currencyfreaks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ffo32167/currencyconverter/internal"
)

type Currencyfreaks struct {
	connStr    string
	currencies string
	client     *http.Client
}

type CurrencyfreaksResponse []currFreaks

type currFreaks struct {
	Date  string            `json:"date"`
	Base  string            `json:"base"`
	Rates map[string]string `json:"rates"`
}

func New(connStr, currencies string, ctxTimeout time.Duration) Currencyfreaks {
	return Currencyfreaks{
		connStr:    connStr,
		currencies: currencies,
		client:     &http.Client{Timeout: ctxTimeout * time.Millisecond}}
}

func (c Currencyfreaks) Rates() ([]internal.Rate, error) {
	resp, err := c.client.Get(c.connStr)

	if err != nil {
		return nil, fmt.Errorf("cant connect with currencyfreaks: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("got wrong response status from currencyfreaks")
	}
	cfr, err := readRates(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read rates: %w", err)
	}
	return cfr.toDomain(c.currencies)
}

func readRates(buf io.Reader) (currFreaks, error) {
	body, err := ioutil.ReadAll(strings.NewReader(`{"date":"2021-11-07 00:03:00+00","base":"USD","rates":{"FJD":"2.085","MATIC":"0.5310110450297366","MXN":"20.3412","STD":"20956.440504","SCR":"13.994033","CDF":"2010.0","BBD":"2.0","HNL":"24.18","UGX":"3552.896508","ZAR":"15.0435","STN":"21.550021","CUC":"1.0","BSD":"1.0","SDG":"440.498","IQD":"1460.0","CUP":"24.3317","GMD":"52.0","TWD":"27.8451","ZRX":"0.7994173846100961","RSD":"101.948123","BSV":"0.005910369687556962","BCH":"0.0017007381203442294","MYR":"4.16","OMG":"0.059897455556087977","FKP":"0.740851","BAND":"0.10740100312536918","XOF":"567.854401","BTC":"0.000016249777276490204","UYU":"43.990579","CVC":"2.0065252200154906","CVE":"95.75","OMR":"0.384996","KES":"111.6","SEK":"8.789025","BTN":"74.28728","GNF":"9550.0","MZN":"63.849999","SVC":"8.749973","ARS":"99.905","QAR":"3.641","IRR":"42025.3","ANKR":"7.765783955890347","SUSHI":"0.08443093549476528","XPD":"0.00049009","ALGO":"0.5445584991967762","THB":"33.257737","UZS":"10700.0","XPF":"103.304111","WBTC":"0.000018","BDT":"85.773927","LYD":"4.555","KWD":"0.301955","XPT":"0.0009635","RUB":"71.1486","ISK":"129.84","MANA":"0.35865355716184527","MKD":"53.288654","DZD":"137.809","PAB":"1.0","SGD":"1.350622","NMR":"0.02200561363203753","JEP":"0.740851","MKR":"0.0003298198069063344","KGS":"84.798902","ZEC":"0.006259389083625439","REN":"1.0563566259969366","REP":"0.03934684241589612","XAF":"567.854401","ADA":"0.4988899698171569","XAG":"0.04139079","STORJ":"0.6006547136378653","CHF":"0.912209","HRK":"6.502","DJF":"178.025","PAX":"1.0","DOGE":"3.813886","TZS":"2301.0","VND":"22642.524891","XAU":"0.00055011","AUD":"1.352271","KHR":"4072.0","IDR":"14319.85","KYD":"0.833348","XRP":"0.867914","BWP":"11.447608","SHP":"0.740851","TJS":"11.269443","AED":"3.673","RWF":"1005.0","DKK":"6.4298","BGN":"1.691345","UMA":"0.07951653944020357","MMK":"1807.978797","NOK":"8.57586","SYP":"2512.53","ZWL":"322.000239","LKR":"201.497397","CZK":"21.8196","XCD":"2.70255","HTG":"98.754554","BHD":"0.377008","CGLD":"0.15776852202448569","KZT":"429.575065","SZL":"15.09","YER":"250.35002","GRT":"0.9428625306430323","AFN":"91.000002","AWG":"1.8","NPR":"118.859851","UNI":"0.03969994779456865","AAVE":"0.0031403484530643524","MNT":"2851.992224","GBP":"0.740851","BYN":"2.451112","HUF":"310.35","BYR":"24511.120000000003","GBX":"19.56123752995559","YFI":"0.000029513072667972914","BIF":"1997.0","XTZ":"0.15538566722605507","XDR":"0.709989","EOS":"0.22484541877459246","BZD":"2.01565","MOP":"8.019334","NAD":"15.09","SKL":"2.6246719160104988","PEN":"4.0175","WST":"2.568092","TMT":"3.5","CLF":"0.029402","GTQ":"7.741234","CLP":"811.4","DNT":"5.253880647593329","TND":"2.8395","COMP":"0.00281610813855252","SLL":"10863.35017","DOP":"56.6","KMF":"426.299814","GEL":"3.16","MAD":"9.0845","TOP":"2.247048","AZN":"1.700805","PGK":"3.53","CNH":"6.3988","UAH":"26.101639","ERN":"15.00062","KNC":"0.5233410090014653","MRO":"356.999828","CNY":"6.3989","ATOM":"0.027604863977032757","MRU":"36.1801","BMD":"1.0","PHP":"50.342996","SNX":"0.0985877307569073","PYG":"6889.520368","JMD":"155.241804","COP":"3868.506465","USD":"1.0","DAI":"0.999455","GGP":"0.740851","ETB":"47.25","ETC":"0.018928997331011378","SOS":"586.0","VEF":"248210.0","VUV":"111.224217","LAK":"10360.0","ETH":"0.00022119856442131693","BND":"1.353506","LRC":"0.778240398459084","LRD":"147.649993","ALL":"107.025","VES":"4.41785","ZMW":"17.381604","BNT":"0.22939725870275848","OXT":"2.0143015409406786","DASH":"0.0052304401153835086","ILS":"3.10978","GHS":"6.1","GYD":"209.359475","KPW":"900.004","BOB":"6.904953","MDL":"17.479552","AMD":"477.286397","TRY":"9.6911","LBP":"1527.84889","JOD":"0.709","GUSD":"1.0","HKD":"7.78395","EUR":"0.865688","LSL":"15.09","CAD":"1.245408","MUR":"43.15","IMP":"0.740851","GIP":"0.740851","RON":"4.2787","NGN":"410.52","CRC":"640.398455","PKR":"170.15","ANG":"1.802173","LTC":"0.005056250790039186","USDC":"1.0","SRD":"21.502","SAR":"3.750973","TTD":"6.7912","CRV":"0.24275085266237","NU":"1.0736525660296328","MVR":"15.45","INR":"74.19115","KRW":"1181.5","JPY":"113.415","AOA":"597.0","PLN":"3.97525","SBD":"8.022495","XLM":"2.7828642351854223","LINK":"0.031196838637159864","MWK":"815.0","MGA":"3967.5","FIL":"0.016165352154114","BAL":"0.03820336453211098","BAM":"1.695716","EGP":"15.720407","SSP":"130.26","BAT":"0.9886835283346812","NIO":"35.225","NZD":"1.405879","ETH2":"0.00022119856442131693","BUSD":"1.0","BRL":"5.5433"}}`))
	if err != nil {
		return currFreaks{}, fmt.Errorf("cant read body of currencyfreaks response: %w", err)
	}

	var cfr currFreaks
	if err := json.Unmarshal([]byte(body), &cfr); err != nil {
		return currFreaks{}, fmt.Errorf("cant unmarshal data from currencyfreaks: %w", err)
	}

	if len(cfr.Rates) == 0 {
		return currFreaks{}, errors.New("cant get rates from currencyfreaks")
	}
	return cfr, nil
}

func (cfr currFreaks) toDomain(currencies string) ([]internal.Rate, error) {
	var rates []internal.Rate
	date, err := time.Parse("2006-01-02 15:04:05+00", cfr.Date)
	if err != nil {
		return nil, fmt.Errorf("cant parse date from currencyfreaks: %w", err)
	}
	/*	rates = append(rates, internal.Rate{
			RateDate: date,
			CurrCode: cfr.Base,
			Rate:     1,
		})
	*/for key, val := range cfr.Rates {
		if strings.Contains(currencies, key) {
			rate, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return nil, fmt.Errorf("cant parse rate value from currencyfreaks: %w", err)
			}
			rates = append(rates, internal.Rate{RateDate: date, CurrCode: key, Rate: rate})
		}
	}
	return rates, nil
}
