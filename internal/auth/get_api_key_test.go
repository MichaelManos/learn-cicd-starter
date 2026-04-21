package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyEmptyAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "fakesite", nil)
	req.Header.Set("Authorization", "")
	_, err := GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("Error did not fire")
	}
}

func TestGetAPIKeySuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "fakesite", nil)
	req.Header.Set("Authorization", "ApiKey timbuktu")
	str, _ := GetAPIKey(req.Header)
	if str != "timbuktu" {
		t.Fatalf("%s does not equal timbuktu", str)
	}
}

func TestGetAPIKeyMalformed(t *testing.T) {
	req, _ := http.NewRequest("GET", "fakesite", nil)
	req.Header.Set("Authorization", "wrongword timbuktu")
	_, err := GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("should have fired malformed error")
	}
}
