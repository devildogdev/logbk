package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

func newEntry() error {
	return nil
}

func handleTilde(p string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, filepath.Base(p))
}

func main() {
	app := viper.New()
	app.SetDefault("journal_path", "~/journal/")
	app.SetConfigType("json")
	app.SetConfigName("settings.json")
	app.AddConfigPath(".")
	err := app.ReadInConfig()
	if err != nil {
		panic(err)
	}
	jpath := app.GetString("journal_path")
	if strings.HasPrefix(jpath, "~") {
		jpath = handleTilde(jpath)
	}
	now := time.Now()
	year := strconv.Itoa(now.Year())
	month := strconv.Itoa(int(now.Month()))
	day := strconv.Itoa(now.Day())
	hour := strconv.Itoa(now.Hour())
	min := strconv.Itoa(now.Minute())
	if len(month) < 2 {
		month = fmt.Sprintf("0%s", month)
	}
	if len(day) < 2 {
		day = fmt.Sprintf("0%s", day)
	}
	ep := fmt.Sprintf("%s/%s/%s/%s%s", jpath, year, month, day, ".md")
	fmt.Println(ep)
	nowEntry := fmt.Sprintf("# %s:%s", hour, min)
	os.WriteFile(ep, []byte(nowEntry), os.ModeAppend)
	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		fmt.Println("$EDITOR not set. Using nvim.")
		editor = "nvim"
	}
	bin, err := exec.LookPath(editor)
	if err != nil {
		panic(err)
	}
	nvimArgs := []string{"nvim", ep}
	env := os.Environ()
	syscall.Exec(bin, nvimArgs, env)
}
