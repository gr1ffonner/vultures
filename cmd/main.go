package main

import (
	"context"
	"log"

	config "github.com/gr1ffonner/vultures/internal"
	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func main() {
	ctx := context.Background()

	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := vault.New(
		vault.WithAddress(cfg.Addr),
	)
	if err != nil {
		log.Fatal(err)
	}

	var loginRequest schema.LdapLoginRequest
	loginRequest.Password = cfg.Password
	resp, err := client.Auth.LdapLogin(ctx, cfg.Login, loginRequest)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.SetToken(resp.Auth.ClientToken); err != nil {
		log.Fatal(err)
	}

	value, err := client.Read(ctx, "v1/dev/data/fiscalization/xi")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(value.Data)
}
