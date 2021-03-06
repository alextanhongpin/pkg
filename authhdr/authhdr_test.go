package authhdr_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/alextanhongpin/pkg/authhdr"
)

func Example() {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	req.Header.Add("Authorization", "Bearer token...")

	hdr := authhdr.New()
	err = hdr.Extract(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hdr)

	switch {
	case hdr.BearerIs(authhdr.Bearer):
		log.Println("is bearer")
		// verifyBearer(hdr.Token())
	case hdr.BearerIs(authhdr.Basic):
		log.Println("is basic")
		// verifyBasic(hdr.Token())
	case hdr.BearerIs("custom"):
		// verifyCustom(hdr.Token())
	default:

	}
}

func TestAuthHeader(t *testing.T) {
	var (
		bearer = "Bearer"
		token  = "token..."
	)
	req, err := http.NewRequest("GET", "http://google.com", nil)
	req.Header.Add("Authorization", bearer+" "+token)

	hdr := authhdr.New()
	err = hdr.Extract(req)
	if err != nil {
		t.Fatal(err)
	}
	if token != hdr.Token() {
		t.Fatalf("expected %s, got %s", token, hdr.Token())
	}
	if bearer != hdr.Bearer() {
		t.Fatalf("expected %s, got %s", bearer, hdr.Bearer())
	}
	if got := hdr.BearerIs(authhdr.Bearer); !got {
		t.Fatalf("expected true, got %t", got)
	}
}
