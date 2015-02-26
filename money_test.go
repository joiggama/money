package money

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	defaults := defaults()

	if defaults["currency"].(string) != "usd" {
		t.Error("Expected default currency to be usd")
	}

	if defaults["no_cents"].(bool) != false {
		t.Error("Expected default no_cents to be false")
	}

	if defaults["with_currency"].(bool) != false {
		t.Error("Expected default with_currency to be false")
	}
}
