package main

import (
	"fmt"
	"github.com/kdhruv3/Discord-Ping/config"
	"github.com/kdhruv3/Discord-Ping/bot"
)

func main(){
	err :=config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make (chan struct {})
	return
}
