package rate

import "net/http"

const (
	ECB_URL        = "https://www.ecb.europa.eu"
	ECB_RATES_PATH = "/stats/eurofxref/eurofxref-daily.xml"
)

type IEcbHttpClient interface {
	GetRates() (*http.Response, error)
}

type EcbHttpClient struct {
	HOST string
}

func MakeEcbHttpClient(host string) *EcbHttpClient {
	return &EcbHttpClient{HOST: host}
}

func (c *EcbHttpClient) GetRates() (*http.Response, error) {
	return http.DefaultClient.Get(c.HOST + ECB_RATES_PATH)
}
