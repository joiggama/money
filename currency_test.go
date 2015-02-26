package money

import (
	"testing"
)

func TestGet(t *testing.T) {
	if get("symbol", "usd").(string) != "$" {
		t.Error(`Expected money.get("symbol", "usd") to return $`)
	}

	if get("symbol", "jpy").(string) != "¥" {
		t.Error(`Expected money.get("symbol", "jpy") to return ¥`)
	}
}

func TestInit(t *testing.T) {
	if currencies == nil {
		t.Error("Expected currencies to be initialized.")
	}
}
