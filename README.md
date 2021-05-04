# goOpenweathermap

func myWeather(APIWeather string, cityID string) string {

  var resultStr string

  n, err := NewWeatherAPI(APIWeather, cityID)
  if err != nil {
    fmt.Println(err)
  }
  //log.Println(n)

resultStr = n.ResultMsg.List[0].Weather[0].Description
  return resultStr
}
