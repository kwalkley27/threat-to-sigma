package feeds

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kwalkley27/threat-to-sigma/config"
)

func TestRetrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "; Spamhaus DROP List")
		fmt.Fprintln(w, "1.2.3.4/24 ; SBL1234")
		fmt.Fprintln(w, "5.6.7.8/24 ; SBL5678")
	}))
	defer server.Close()

	cfg := &config.Config{
		SpamhausFeedURL: server.URL,
		FeedLimit:       2,
	}

	iocs, err := Retrieve(cfg)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(iocs) != 2 {
		t.Fatalf("expected 2 IOCs, got %d", len(iocs))
	}

	if iocs[0] != "1.2.3.4/24" {
		t.Errorf("expected IOC 1.2.3.4/24, got %s", iocs[0])
	}

	if iocs[1] != "5.6.7.8/24" {
		t.Errorf("expected IOC 5.6.7.8/24, got %s", iocs[1])
	}
}
