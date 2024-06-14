package main

import (
	"fmt"

	"github.com/jaddek/ecb/rate"
)

func main() {
	httpClient := rate.MakeEcbHttpClient(rate.ECB_URL)
	body := rate.GetEcbRates(httpClient)
	envelope := rate.MakeEnvelope(body)

	fmt.Println(string(envelope.GetEnvelopeAsJson()))
}
