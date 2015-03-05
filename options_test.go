package money

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	defaults := defaults()

	if defaults["currency"].(string) != "USD" {
		t.Error("Expected default currency to be USD")
	}

	if defaults["with_cents"].(bool) != true {
		t.Error("Expected default with_cents to be true")
	}

	if defaults["with_currency"].(bool) != false {
		t.Error("Expected default with_currency to be false")
	}
}

func TestOverride(t *testing.T) {
	options := defaults()

	options = override(options, Options{"currency": "CAD"})

	if options["currency"].(string) != "CAD" {
		t.Error("Expected currency to be overriden to CAD")
	}

	if options["with_currency"].(bool) != false {
		t.Error("Expected with_currency not to be overriden")
	}
}
