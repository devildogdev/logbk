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

func newEntryPath(root string) (string, error) {
	if strings.HasPrefix(root, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		root = filepath.Join(homeDir, filepath.Base(root))
	}
	if !filepath.IsAbs(root) {
		return "", fmt.Errorf("%s is not an absolute path!", root)
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
	return fmt.Sprintf("%s/%s/%s/%s%s", root, y, m, d, ".md"), nil
}

func openFileInEditor(path string) error {
	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		fmt.Println("$EDITOR not set. Using nvim.")
		editor = "nvim"
	}
	bin, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	nvimArgs := []string{"nvim", path}
	env := os.Environ()
	syscall.Exec(bin, nvimArgs, env)
	return nil
}

func main() {
	conf := getConfig()
	ep, err := newEntryPath(conf.GetString("journal_path"))
	if err != nil {
		panic(err)
	}
	openFileInEditor(ep)
}
