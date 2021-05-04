package main

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
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

type ResultMsg struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []List `json:"list"`
	City    City
}

type List struct {
	Dt         float64   `json:"dt"`
	Main       Main      `json:"main"`
	Weather    []Weather `json:"weather"`
	Clouds     Clouds    `json:"clouds"`
	Wind       Wind      `json:"wind"`
	Visibility int       `json:"visibility"`
	Pop        float64   `json:"pop"`

	DateTxt string `json:"dt_txt"`
}

type Main struct {
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	Pressure    int64   `json:"pressure"`
	SeaLevel    int64   `json:"sea_level"`
	GroundLevel int64   `json:"grnd_level"`
	Humidity    int64   `json:"humidity"`
	TempKf      float64 `json:"temp_kf"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CityCoord  CityCoord
	Country    string  `json:"country"`
	Population float64 `json:"population"`
	timezone   float64 `json:"timezone"`
	sunrise    float64 `json:"sunrise"`
	sunset     float64 `json:"sunset"`
}

type CityCoord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type WeatherAPI struct {
	Token       string     `json:"token"`
	Debug       bool       `json:"debug"`
	ResultMsg   ResultMsg  `json:"-"`
	Client      HttpClient `json:"-"`
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
	response, err := weather.MakeRequest(cityID, nil) //for Khabarovsk, test "2022890"
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
