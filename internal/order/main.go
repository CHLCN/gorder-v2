package main

import (
	"log"

	"github.com/CHLCN/gorder-v2/common/config"
	"github.com/spf13/viper"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("%v", viper.Get("order"))

}
