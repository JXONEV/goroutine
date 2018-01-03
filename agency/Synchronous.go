package agency

func Syn() {
    customers := GetCustomerDetails()
    destinations := GetRecommendedDestinations(customers)
	  var Infos [10]Info
	  for i, d := range destinations {
		    qoute := GetQuote(d)
		    weather := GetWeatherForecast(d)
		    Infos[i] = Info{Des:d, Quo:qoute, Wea:weather}
	  }
}
