package main

import (
	"github.com/dice2005x3005/RSS/internal/config"
)

func main(){
	conf, err := config.Read()
	if err != nil {
		return
	}
	conf.SetUser("Marco")
	conf, err := config.Read()
	if err != nil {
		return
	}
	fmt.Printf("%v", conf)
}