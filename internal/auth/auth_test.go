package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	key, err := GetAPIKey(http.Header{})
	if key != "" || err != ErrNoAuthHeaderIncluded {
		t.Errorf("empty header should return blank key")
	}
}
