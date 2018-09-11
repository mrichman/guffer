package main

import "testing"

// Test loading twitter auth keys from file
func TestTwitterAuthKeysToml(t *testing.T) {
	// Load twitter auth keys from file
	var auth TwitterAuthKeys
	auth.loadFromTomlFile("./testfiles/auth.toml")
	// Compare loaded auth data with expected auth data
	expectedAuth := TwitterAuthKeys{
		ConsumerKey:       "a",
		ConsumerSecret:    "b",
		AccessToken:       "c",
		AccessTokenSecret: "d",
	}
	if auth != expectedAuth {
		t.Fail()
	}
}
