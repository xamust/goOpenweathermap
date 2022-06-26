package openweathermap

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const APIEndpoint = "https://api.openweathermap.org/data/2.5/forecast?id=%s&APPID=%s"

type APIResponse struct {
	Ok          bool
	Result      json.RawMessage
	ErrorCode   int
	Description string
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherAPI struct {
	Token       string
	Debug       bool
	ResultMsg   ResultMsg
	Client      HttpClient
	apiEndpoint string
}

type Error struct {
	Code    int
	Message string
}

func NewWeatherAPI(token string, cityID string) (*WeatherAPI, error) {
	return NewWeatherAPIWithToken(token, cityID, APIEndpoint, &http.Client{})
}

func NewWeatherAPIWithToken(token, cityID, apiEndpoint string, client HttpClient) (*WeatherAPI, error) {
	weather := &WeatherAPI{
		Token:       token,
		Debug:       false,
		Client:      client,
		apiEndpoint: apiEndpoint,
	}
	myWeather, error := weather.GetWeather(cityID)
	if error != nil {
		return nil, error
	}
	weather.ResultMsg = myWeather
	return weather, nil
}

func (weather *WeatherAPI) GetWeather(cityID string) (ResultMsg, error) {
	response, err := weather.MakeRequest(cityID, nil)
	if err != nil {
		return ResultMsg{}, err //APIResponse
	}
	return response, err
}

func (weather *WeatherAPI) MakeRequest(cityID string, params url.Values) (ResultMsg, error) {
	method := fmt.Sprintf(weather.apiEndpoint, cityID, weather.Token)
	req, err := http.NewRequest("POST", method, strings.NewReader(params.Encode()))
	if err != nil {
		return ResultMsg{}, err //ResultMsg
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := weather.Client.Do(req)
	if err != nil {
		return ResultMsg{}, err
	}
	defer resp.Body.Close()

	var apiResp ResultMsg

	bytes, err := weather.decodeAPIResponse(resp.Body, &apiResp)
	if err != nil {
		return apiResp, err
	}

	if weather.Debug {
		log.Printf("%s resp: %s", cityID, bytes)
	}
	/*
		if !apiResp.Ok {
			parameters := ResponseParameters{}
			if apiResp.Parameters != nil {
				parameters = *apiResp.Parameters
			}
			return apiResp, err
		}
	*/
	return apiResp, nil
}

func (weather *WeatherAPI) decodeAPIResponse(responseBody io.Reader, resp *ResultMsg) (_ []byte, err error) {
	if !weather.Debug {
		dec := json.NewDecoder(responseBody)
		err = dec.Decode(resp)
		return
	}

	// if debug, read reponse body
	data, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, resp)
	if err != nil {
		return
	}

	return data, nil
}
