package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func getConfig() *viper.Viper {
	app := viper.New()
	app.SetDefault("journal_path", "~/journal/")
	app.SetConfigType("json")
	app.SetConfigName("settings.json")
	app.AddConfigPath(".")
	err := app.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func getEntryPath() {}

func main() {
	conf := getConfig()
	jpath := conf.GetString("journal_path")
	if strings.HasPrefix(jpath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		jpath = filepath.Join(homeDir, filepath.Base(jpath))
	}
	fmt.Printf("journal_path: %s\n", jpath)
	now := time.Now()
	yyyy := now.Year()
	mm := now.Month()
	dd := now.Day()
	if mm < 10 {
		fmt.Printf("generated_path: %s/%d/0%d/%d%s\n", jpath, yyyy, mm, dd, ".md")
	} else {
		fmt.Printf("generated_path: %s/%d/%d/%d%s\n", jpath, yyyy, mm, dd, ".md")
	}
}
