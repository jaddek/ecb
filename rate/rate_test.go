package rate

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expected     = `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope xmlns:gesmes="http://www.gesmes.org/xml/2002-08-01" xmlns="http://www.ecb.int/vocabulary/2002-08-01/eurofxref"><gesmes:subject>Referencerates</gesmes:subject><gesmes:Sender><gesmes:name>EuropeanCentralBank</gesmes:name></gesmes:Sender><Cube><Cube time='2024-06-13'><Cube currency='USD' rate='1.0784'/><Cube currency='JPY' rate='169.58'/><Cube currency='BGN' rate='1.9558'/><Cube currency='CZK' rate='24.699'/><Cube currency='DKK' rate='7.4593'/><Cube currency='GBP' rate='0.84468'/><Cube currency='HUF' rate='396.48'/><Cube currency='PLN' rate='4.3473'/><Cube currency='RON' rate='4.9773'/><Cube currency='SEK' rate='11.2210'/><Cube currency='CHF' rate='0.9668'/><Cube currency='ISK' rate='149.30'/><Cube currency='NOK' rate='11.4315'/><Cube currency='TRY' rate='34.8311'/><Cube currency='AUD' rate='1.6232'/><Cube currency='BRL' rate='5.8261'/><Cube currency='CAD' rate='1.4823'/><Cube currency='CNY' rate='7.8211'/><Cube currency='HKD' rate='8.4224'/><Cube currency='IDR' rate='17527.83'/><Cube currency='ILS' rate='4.0108'/><Cube currency='INR' rate='90.1120'/><Cube currency='KRW' rate='1482.31'/><Cube currency='MXN' rate='20.1654'/><Cube currency='MYR' rate='5.0766'/><Cube currency='NZD' rate='1.7477'/><Cube currency='PHP' rate='63.173'/><Cube currency='SGD' rate='1.4557'/><Cube currency='THB' rate='39.593'/><Cube currency='ZAR' rate='19.8385'/></Cube></Cube></gesmes:Envelope>`
	expectedJson = `{"data":{"rates":{"date":"2024-06-13","pairs":[{"currency":"USD","rate":"1.0784"},{"currency":"JPY","rate":"169.58"},{"currency":"BGN","rate":"1.9558"},{"currency":"CZK","rate":"24.699"},{"currency":"DKK","rate":"7.4593"},{"currency":"GBP","rate":"0.84468"},{"currency":"HUF","rate":"396.48"},{"currency":"PLN","rate":"4.3473"},{"currency":"RON","rate":"4.9773"},{"currency":"SEK","rate":"11.2210"},{"currency":"CHF","rate":"0.9668"},{"currency":"ISK","rate":"149.30"},{"currency":"NOK","rate":"11.4315"},{"currency":"TRY","rate":"34.8311"},{"currency":"AUD","rate":"1.6232"},{"currency":"BRL","rate":"5.8261"},{"currency":"CAD","rate":"1.4823"},{"currency":"CNY","rate":"7.8211"},{"currency":"HKD","rate":"8.4224"},{"currency":"IDR","rate":"17527.83"},{"currency":"ILS","rate":"4.0108"},{"currency":"INR","rate":"90.1120"},{"currency":"KRW","rate":"1482.31"},{"currency":"MXN","rate":"20.1654"},{"currency":"MYR","rate":"5.0766"},{"currency":"NZD","rate":"1.7477"},{"currency":"PHP","rate":"63.173"},{"currency":"SGD","rate":"1.4557"},{"currency":"THB","rate":"39.593"},{"currency":"ZAR","rate":"19.8385"}]}}}`
)

func TestGetEcbRates(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()

	httpClient := MakeEcbHttpClient(svr.URL)

	content := string(GetEcbRates(httpClient))

	assert.Equal(t, content, expected)
}

func TestMakeEnvolope(t *testing.T) {
	envelope := MakeEnvelope([]byte(expected))

	assert.Equal(t, "2024-06-13", envelope.GetDate())

	currency, _ := envelope.GetRateValueByCurrency("USD")
	assert.NotEmpty(t, currency)

	currencies := envelope.GetCurrencies()
	assert.NotEmpty(t, currencies)
	assert.Equal(t, 30, len(currencies))
}

func TestEnvelopeGetEnvelopeAsJson(t *testing.T) {
	envelope := MakeEnvelope([]byte(expected))

	assert.JSONEq(t, expectedJson, string(envelope.GetEnvelopeAsJson()))
}
