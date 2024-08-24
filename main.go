package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func getConfig() map[string]any {
	v := viper.New()
	v.SetDefault("journal_path", "~/journal")
	v.SetConfigFile("./settings.json")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return v.AllSettings()
}

func newEntry() {}

func main() {
	// conf := getConfig()
	viper.SetDefault("journal_path", "~/journal")
	viper.SetConfigFile("./settings.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	conf := viper.AllSettings()
	fmt.Println(conf["journal_path"])
}
