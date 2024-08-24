package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("journal_path", "~/journal")
	viper.SetConfigFile("./settings.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf := viper.AllSettings()
	fmt.Println(conf["journal_path"])
}
