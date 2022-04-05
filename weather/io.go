package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strzybug/utils"
)

// Request models OpenWeatherMap OneCall API
type Request struct {
	Latitude     string
	Longitude    string
	LanguageCode string
	ApiKey       string
}

func (c Request) Url() string {
	return fmt.Sprintf("https://api.openweathermap.org/data/2.5/onecall?exclude=current,minutely,hourly&units=metric&lat=%s&lon=%s&appid=%s&lang=pl",
		c.Latitude, c.Longitude, c.ApiKey)
}

func (c Request) Run() (res Response, err error) {
	var (
		httpResp *http.Response
		bytes []byte
	)

	if httpResp, err = http.Get(c.Url()); err == nil {
		if bytes, err = ioutil.ReadAll(httpResp.Body); err == nil {
			err = json.Unmarshal(bytes, &res)
		}
	}
	return res, utils.Wrap(err)
}

func FromFile(filename string) (res Response, err error) {
	var file []byte
	if file, err = ioutil.ReadFile(filename); err == nil {
		err = json.Unmarshal(file, &res)
	}
	return res, utils.Wrap(err)
}

func (r Response) ToFile(filename string) (err error) {
	var (
		file *os.File
		bytes []byte
	)
	if file, err = os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644); err == nil {
		defer file.Close()
		if bytes, err = json.MarshalIndent(r, "", "  "); err == nil {
			_, err = file.Write(bytes)
		}
	}
	return utils.Wrap(err)
}
