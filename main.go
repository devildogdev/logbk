package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
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

func newEntryPath(root string) string {
	if strings.HasPrefix(root, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		root = filepath.Join(homeDir, filepath.Base(root))
	}
	now := time.Now()
	y := strconv.Itoa(now.Year())
	m := strconv.Itoa(int(now.Month()))
	d := strconv.Itoa(now.Day())
	if len(m) < 2 {
		m = fmt.Sprintf("0%s", m)
	}
	if len(d) < 2 {
		d = fmt.Sprintf("0%s", d)
	}
	return fmt.Sprintf("%s/%s/%s/%s%s", root, y, m, d, ".md")
}

func main() {
	conf := getConfig()
	ep := newEntryPath(conf.GetString("journal_path"))
	fmt.Println(ep)
}
