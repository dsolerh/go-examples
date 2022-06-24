package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	pflag.StringP("name", "n", "Mike", "Name parameter")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	migdir, err := checkdir(cwd, "migrations")
	if err != nil {
		log.Fatal(err)
	}
}
