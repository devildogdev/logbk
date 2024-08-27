package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

var editorCmds = []string{
	"vim",
	"nvim",
}

func twoDigitString(n int) string {
	s := strconv.Itoa(n)
	if len(s) < 2 {
		s = fmt.Sprintf("0%s", s)
	}
	return s
}

func addTimestamp(p string, ts string) error {
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = f.WriteString(ts)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func checkEntryExists(p string) {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		d := filepath.Dir(p)
		os.MkdirAll(d, 0755)
	}
}

func openWithEditor(fp string) error {
	editor, ok := os.LookupEnv("EDITOR")
	if !ok {
		fmt.Println("$EDITOR not set. Using vim.")
		editor = "vim"
	}
	if !slices.Contains(editorCmds, editor) {
		fmt.Println("Your editor is not supported. Using vim.")
		editor = "vim"
	}
	bin, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	args := []string{editor, "+", fp}
	err = syscall.Exec(bin, args, os.Environ())
	if err != nil {
		return err
	}
	return nil
}

func newEntry(path string) error {
	now := time.Now()
	ep := fmt.Sprintf(
		"%s/%s/%s/%s%s",
		path,
		twoDigitString(now.Year()),
		twoDigitString(int(now.Month())),
		twoDigitString(now.Day()),
		".md",
	)
	checkEntryExists(ep)
	ts := fmt.Sprintf(
		"\n\n# %s:%s\n\n\n",
		twoDigitString(now.Hour()),
		twoDigitString(now.Minute()),
	)
	addTimestamp(ep, ts)
	err := openWithEditor(ep)
	if err != nil {
		return err
	}
	return nil
}

func handleTilde(jp string) string {
	h, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(
		h,
		filepath.Base(jp),
	)
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
	err = newEntry(jpath)
	if err != nil {
		panic(err)
	}
}
