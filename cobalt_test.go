package main

import (
	"testing"

	"github.com/lostdusty/gobalt"
)

//These tests allows me to test the application directly w/o the need to run `go run . <sub> [flags]`.

func TestInstacesListing(t *testing.T) {
	t.Log("testing instance listing")
	err := getInstances(false)
	if err != nil {
		t.Fatalf("failed to fetch custom instaces: %v", err)
	}
}

func TestCobaltHealth(t *testing.T) {
	err := checkStatus(gobalt.CobaltApi, true)
	if err != nil {
		t.Fatalf("failed to get cobalt instance health: %v", err)
	}
}
