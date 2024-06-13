package converter

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

type Envelope struct {
	Cube struct {
		Cube struct {
			Date  string `xml:"time,attr" json:"date"`
			Rates []struct {
				Currency string `xml:"currency,attr" json:"currency"`
				Rate     string `xml:"rate,attr" json:"rate"`
			} `xml:"Cube" json:"pairs"`
		} `xml:"Cube" json:"rates"`
	} `xml:"Cube" json:"data"`
}

func Convert(host string) []byte {
	resp, err := http.Get(host)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := &Envelope{}
	err2 := xml.Unmarshal(body, data)
	if nil != err2 {
		log.Fatal("Error unmarshalling from XML", err)
	}

	result, err := json.Marshal(data)
	if nil != err {
		log.Fatal("Error marshalling to JSON", err)
	}

	return result
}
