# goOpenweathermap

func myWeather(APIWeather string, cityID string) string {
  var resultStr string
  
  n, err := NewWeatherAPI(APIWeather, cityID)
  
  if err != nil {
    log.Println(err)
  }

  resultStr = n.ResultMsg.List[0].Weather[0].Description
  log.Println(resultStr)

  return resultStr
}
