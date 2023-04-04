package mplus_test

import (
	"os"
	"testing"

	"github.com/omniboost/go-mplus"
)

var (
	client *mplus.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	ident := os.Getenv("IDENT")
	secret := os.Getenv("SECRET")
	debug := os.Getenv("DEBUG")

	client = mplus.NewClient(nil)
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
