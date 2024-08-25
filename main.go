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

func main() {
	viper.SetDefault("journal_path", "~/journal/")
	viper.SetConfigFile("./settings.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	jpath := viper.GetString("journal_path")
	if strings.HasPrefix(jpath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		jpath = filepath.Join(homeDir, filepath.Base(jpath))
	}
	fmt.Printf("journal_path: %s\n", jpath)
	// err = filepath.WalkDir(
	// 	jpath,
	// 	func(r string, d fs.DirEntry, err error) error {
	// 		info, err := d.Info()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		fmt.Println(info.Name())
	// 		return nil
	// 	})
	// if err != nil {
	// 	log.Fatal(err)
	// }
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
