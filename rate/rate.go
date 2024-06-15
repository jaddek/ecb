package rate

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
)

type Envelope struct {
	Cube struct {
		Cube struct {
			Date  string `json:"date"  xml:"time,attr"`
			Rates []Rate `json:"pairs" xml:"Cube"`
		} `json:"rates" xml:"Cube"`
	} `json:"data" xml:"Cube"`
}

type Rate struct {
	Currency string `json:"currency" xml:"currency,attr"`
	Rate     string `json:"rate"     xml:"rate,attr"`
}

func (envelope *Envelope) GetDate() string {
	return envelope.Cube.Cube.Date
}

func (envelope *Envelope) GetCurrencies() []string {
	rates := envelope.GetRates()

	currencies := make([]string, 0)
	for _, rate := range rates {
		currencies = append(currencies, rate.Currency)
	}

	return currencies
}

func (envelope *Envelope) GetCurrenciesAsJson() []byte {
	currencies := envelope.GetCurrencies()

	return asJson(currencies)
}

func (envelope *Envelope) GetRatesAsJson() []byte {
	rates, err := json.Marshal(envelope.GetRates())

	if nil != err {
		log.Fatal("Error marshalling to JSON", err)
	}

	return rates
}

func (envelope *Envelope) GetRateObjectByCurrency(currency string) (*Rate, error) {
	rates := envelope.GetRates()

	for _, rate := range rates {
		if rate.Currency == currency {
			return &rate, nil
		}
	}

	return nil, fmt.Errorf("currency %s not found", currency)
}

func (envelope *Envelope) GetRateByCurrencyAsJson(currency string) []byte {
	rate, err := envelope.GetRateObjectByCurrency(currency)

	if err != nil {
		log.Panic(err)
	}

	return asJson(rate)
}

func (envelope *Envelope) GetRateValueByCurrency(currency string) (string, error) {
	rates := envelope.GetRates()

	for _, rate := range rates {
		if rate.Currency == currency {
			return string(rate.Rate), nil
		}
	}

	return "", fmt.Errorf("currency %s not found", currency)
}

func (envelope *Envelope) GetRates() []Rate {
	return envelope.Cube.Cube.Rates
}

func (envelope *Envelope) GetEnvelopeAsJson() []byte {
	return asJson(envelope)
}

func asJson(o interface{}) []byte {
	result, err := json.Marshal(o)
	if nil != err {
		log.Fatal("Error marshalling to JSON", err)
	}

	return result
}

func GetEcbRates(httpClient IEcbHttpClient) []byte {
	resp, err := httpClient.GetRates()

	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func MakeEnvelope(body []byte) *Envelope {
	envelope := &Envelope{}
	err := xml.Unmarshal(body, envelope)

	if nil != err {
		log.Fatal("Error unmarshalling from XML", err)
	}

	return envelope
}
