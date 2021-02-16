package main

import (
	"fmt"

	config "github.com/02amanag/getconfig/configuration"
)

func main() {

	// .env is the eviroment file where all the value of global variable are defined in key value pair
	conf := config.NewConfig(".env")

	// SECRET_KEY is one of the key whose value is fetch and used
	value, err := conf.GetConfig("SECRET_KEY")
	if err != nil {
		panic(err)
	} else {
		// value store the value of that key
		fmt.Println("Value we retrive - ", value)
	}

}
