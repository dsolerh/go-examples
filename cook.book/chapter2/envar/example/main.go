package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dsolerh/examples/cook.book/chapter2/envar"
)

// Config holds the config we
// capture from a json file and env vars
type Config struct {
	Version string `json:"version" required:"true"`
	IsSafe  bool   `json:"is_safe" default:"true"`
	Secret  string `json:"secret"`
}

func main() {
	// create a termporary file to hold
	// an example json file
	tf, err := ioutil.TempFile("", "tmp")
	if err != nil {
		panic(err)
	}
	defer tf.Close()
	defer os.Remove(tf.Name())

	// create a json file to hold our secrets
	secrets := `{
		"secret":"so so secret"
	}`

	if _, err := tf.Write(bytes.NewBufferString(secrets).Bytes()); err != nil {
		panic(err)
	}

	// we can easily set enviroment variables as needed
	if err := os.Setenv("EXAMPLE_VERSION", "1.0.0"); err != nil {
		panic(err)
	}
	if err := os.Setenv("EXAMPLE_ISSAFE", "false"); err != nil {
		panic(err)
	}

	c := Config{}
	if err := envar.LoadConfig(tf.Name(), "EXAMPLE", &c); err != nil {
		panic(err)
	}

	fmt.Println("secret file contains =", secrets)

	// we can also read them
	fmt.Println("EXAMPLE_VERSION =", os.Getenv("EXAMPLE_VERSION"))
	fmt.Println("EXAMPLE_ISSAFE =", os.Getenv("EXAMPLE_ISSAFE"))
	// The final config is a mix of json and environment
	// variables
	fmt.Printf("Final Config: %#v\n", c)
}
