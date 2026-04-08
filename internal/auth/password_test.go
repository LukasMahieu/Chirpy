package auth

import (
	"testing"
)

func TestHashUnHashPassword(t *testing.T) {
	cases := []struct {
		name      string
		password  string
		input     string
		wantMatch bool
	}{
		{name: "correct password", password: "temp", input: "temp", wantMatch: true},
		{name: "wrong password", password: "temp", input: "wrong", wantMatch: false},
		{name: "empty password", password: "", input: "", wantMatch: true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			hash, err := HashPassword(tc.password)
			if err != nil {
				t.Fatalf("unexpected error hashing: %v", err)
			}
			if hash == "" {
				t.Fatalf("Expected non-empty hash")
			}
			match, err := CheckPasswordHash(tc.input, hash)
			if err != nil {
				t.Fatalf("unexpected error checking hash: %v", err)
			}
			if match != tc.wantMatch {
				t.Fatalf("got match=%v want match=%v", match, tc.wantMatch)
			}
		})
	}
}
