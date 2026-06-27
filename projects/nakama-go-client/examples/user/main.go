// Example: authenticate a player with a custom identifier.
//
// Run with:
//
//	go run ./examples/user
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dsolerh/nkclient"
)

func main() {
	client, err := nkclient.NewClient(
		"defaultkey",
		nkclient.DefaultHost,
		nkclient.WithPort(nkclient.DefaultAPIPort),
		nkclient.WithSSL(false),
		nkclient.WithTimeout(5*time.Second),
	)
	if err != nil {
		log.Fatalf("new client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	session, err := client.AuthenticateCustom(ctx, "external-user-id-42", nkclient.AuthenticateRequest{
		Create:   true,
		Username: "ada",
	})
	if errors.Is(err, nkclient.ErrNotImplemented) {
		log.Printf("AuthenticateCustom is a stub — wire up the transport before running for real")
		return
	}

	fmt.Printf(
		"signed in as %s (created=%t, expires=%s)\n",
		session.Username,
		session.Created,
		session.ExpiresAt.Format(time.RFC3339),
	)

	type payload struct {
		Data string `json:"data"`
	}
	type response struct {
		Value int `json:"value"`
	}
	resp, err := nkclient.CallCustomRPC[payload, response](ctx, client, session, "some_rpc", payload{Data: ""})
	fmt.Printf("err: %v\n", err)
	fmt.Printf("resp: %v\n", resp)
}
