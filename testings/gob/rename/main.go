// Checks what happens when a struct is gob-encoded while registered under one
// name and decoded while registered under another name.
//
// gob only uses the registered name for values sent inside an interface, so the
// struct is wrapped in an `any` field.
//
// A type can only be registered once per process (a second RegisterName with a
// different name panics), so the program re-runs itself: one child process
// encodes, another child process decodes.
//
// Run: go run ./gob/rename
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"os/exec"
	"reflect"
)

type Player struct {
	ID    int
	Name  string
	Score float64
}

type Envelope struct {
	Payload any
}

const (
	encodeName = "player.v1"
	decodeName = "player.v2"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "encode":
			encode()
		case "decode":
			decode()
		case "decode-compat":
			decodeCompat()
		}
		return
	}

	fmt.Println("== reflect.TypeOf")
	fmt.Println("Player:  ", reflect.TypeOf(Player{}))
	fmt.Println("PlayerV1:", reflect.TypeOf(PlayerV1{}))

	fmt.Println("\n== same process: register the type under two names")
	sameProcess()

	fmt.Println("\n== separate processes: encode as", encodeName, "-> decode as", decodeName)
	data := run("encode", nil)
	fmt.Printf("encoded %d bytes\n", len(data))
	out := run("decode", data)
	fmt.Print(string(out))

	fmt.Println("\n== fix: legacy type registered under the old name")
	out = run("decode-compat", data)
	fmt.Print(string(out))
}

func encode() {
	gob.RegisterName(encodeName, Player{})
	var buf bytes.Buffer
	env := Envelope{Payload: Player{ID: 7, Name: "ana", Score: 99.5}}
	if err := gob.NewEncoder(&buf).Encode(env); err != nil {
		fmt.Fprintln(os.Stderr, "encode error:", err)
		os.Exit(1)
	}
	os.Stdout.Write(buf.Bytes())
}

func decode() {
	gob.RegisterName(decodeName, Player{})
	var env Envelope
	err := gob.NewDecoder(os.Stdin).Decode(&env)
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Printf("decoded: %#v\n", env.Payload)
}

// PlayerV1 has the same fields as Player but is a separate type, so it can be
// registered under the old name while Player keeps the new one.
type PlayerV1 Player

func decodeCompat() {
	gob.RegisterName(decodeName, Player{})
	gob.RegisterName(encodeName, PlayerV1{})
	var env Envelope
	if err := gob.NewDecoder(os.Stdin).Decode(&env); err != nil {
		fmt.Println("decode error:", err)
		return
	}
	if old, ok := env.Payload.(PlayerV1); ok {
		env.Payload = Player(old)
	}
	fmt.Printf("decoded: %#v\n", env.Payload)
}

func sameProcess() {
	type Local struct{ X int }
	gob.RegisterName("local.a", Local{})
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()
	gob.RegisterName("local.b", Local{})
	fmt.Println("no panic")
}

func run(mode string, stdin []byte) []byte {
	cmd := exec.Command(os.Args[0], mode)
	cmd.Stdin = bytes.NewReader(stdin)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(mode, "failed:", err)
		os.Exit(1)
	}
	return out
}
