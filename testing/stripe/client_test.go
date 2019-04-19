package stripe_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/zsanders16/go_tutorials/testing/stripe"
)

var (
	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "key", "", "Your test secrect key for the Stripe API. If present integration test will be run using this key.")
}

func TestClient_Customer(t *testing.T) {
	if apiKey == "" {
		t.Fatal("No API key provided.")
	}

	c := stripe.Client{
		Key: apiKey,
	}
	tok := "tok_amex"
	email := "test@test.com"
	cus, err := c.Customer(tok, email)
	if err != nil {
		t.Errorf("Customer() err = %v; want %v", err, nil)
	}

	if cus == nil {
		t.Fatalf("Customer() = nil; want non-nill value")
	}

	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.ID, "cus_")
	}
	if !strings.HasPrefix(cus.DefaultSource, "card_") {
		t.Errorf("Customer() DefaultSource = %s; want prefix %q", cus.DefaultSource, "card_")
	}
	if cus.Email != email {
		t.Errorf("Customer() Email = %s; want %s", cus.Email, email)
	}

}
