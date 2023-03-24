package mplus_test

import (
	"os"
	"testing"

	netsuite "github.com/omniboost/go-netsuite-soap"
)

var (
	client *netsuite.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	ident := os.Getenv("IDENT")
	secret := os.Getenv("SECRET")
	debug := os.Getenv("DEBUG")

	client = netsuite.NewClient(nil)
	client.SetIdent(ident)
	client.SetSecret(secret)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		client.SetBaseURL(baseURL)
	}

	m.Run()
}
