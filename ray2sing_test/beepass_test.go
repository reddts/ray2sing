package ray2sing_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/reddts/ray2sing/ray2sing"
)

func TestBeePass(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"server":"beacomf.xyz","server_port":"8080","password":"nfzmfcBTcsj287NxNgMZDu","method":"chacha20-ietf-poly1305","name":"BeePass"}`)
	}))
	defer ts.Close()

	url := strings.Replace(ts.URL, "http://", "ssconf+http://", 1) + "/config.json#BeePass"

	// Define the expected JSON structure
	expectedJSON := `{
		"outbounds": [
			{
				"type": "shadowsocks",
				"tag": "BeePass § 0",
				"server": "beacomf.xyz",
				"server_port": 8080,
				"method": "chacha20-ietf-poly1305",
				"password": "nfzmfcBTcsj287NxNgMZDu"
			}
		]
	}`
	ray2sing.CheckUrlAndJson(url, expectedJSON, t)
}
