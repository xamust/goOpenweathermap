# goOpenweathermap

func myWeather(APIWeather string, cityID string) string {

  var resultStr string



  n, err := NewWeatherAPI(APIWeather, cityID)
  if err != nil {
    fmt.Println(err)
  }
  //log.Println(n)

resultStr = n.ResultMsg.List[0].Weather[0].Description
log.Println(resultStr)
  for k := 0; k < len(n.ResultMsg.List); k++ {

    //fmt.Println(n.ResultMsg.List[k])
  }
  return resultStr
}
