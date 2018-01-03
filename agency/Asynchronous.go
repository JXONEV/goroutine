package agency

func Asy() {
    customer := GetCustomerDetails()
    destinations := GetRecommendedDestinations(customer)
	  var Infos [10]Info
	  quotes := [10]chan Quoting{}
	  weathers := [10]chan Weather{}

	  // create channel Quoting
	  for i := range quotes {
		    quotes[i] = make(chan Quoting)
	  }
	  // create channel Weather
	  for i := range weathers {
		    weathers[i] = make(chan Weather)
	  }

 	  for i, d := range destinations {
		    temp := i
		    d_temp := d
		    go func() {
			      quotes[temp] <- GetQuote(d_temp) // receive Quoting from channel
		    }()
	      go func() {
			      weathers[temp] <- GetWeatherForecast(d_temp) // receive Weather from channel
		    }()
	  }
	  // wait for receive
	  for i, d := range destinations {
			  Infos[i] = Info{Des:d, Quo:<-quotes[i], Wea:<-weathers[i]}
	  }
}
