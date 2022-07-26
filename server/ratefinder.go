package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"gopkg.in/resty.v0"
	"net/http"
	"service1/models"
)

func (srv *Server) GetRate(resp http.ResponseWriter, req *http.Request) {
	temp, err := resty.R().
		SetHeader("X-CoinAPI-Key", "0C3C7B3D-23FC-4541-9994-8A38E24C2752").
		Get("https://rest.coinapi.io/v1/exchangerate/BTC?invert=false")

	if err != nil {
		logrus.Errorf("something went wrong")
		return
	}

	var data models.CoinInfo
	err = json.Unmarshal(temp.Body, &data)
	if err != nil {
		logrus.Errorf("error in unmarshalling data")
		return
	}

	var sendData float64
	for i := range data.RateInfo {
		if data.RateInfo[i].AssetID == "USD" {
			sendData = data.RateInfo[i].Price
		}
	}

	//res, err := http.Get("https://api.coinlayer.com/live?access_key=YOUR_ACCESS_KEY&expan=1")
	err = json.NewEncoder(resp).Encode(sendData)
	if err != nil {
		logrus.Errorf("error in encodeing data")
		return
	}
}

func (srv *Server) Health(resp http.ResponseWriter, req *http.Request) {
	_, err := resp.Write([]byte("Success   comes with a price"))
	if err != nil {
		return
	}
}
