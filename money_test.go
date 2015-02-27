package money

import (
	"testing"
)

func TestDefaults(t *testing.T) {
	defaults := defaults()

	if defaults["currency"].(string) != "usd" {
		t.Error("Expected default currency to be usd")
	}

	if defaults["with_cents"].(bool) != true {
		t.Error("Expected default with_cents to be true")
	}

	if defaults["with_currency"].(bool) != false {
		t.Error("Expected default with_currency to be false")
	}
}

func TestNewDefaults(t *testing.T) {
	currency := New(10)
	expected := "$10.00"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWithCurrency(t *testing.T) {
	currency := New(10, Options{"with_currency": true})
	expected := "$10.00 USD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWithCurrencyWhenCanadian(t *testing.T) {
	currency := New(10, Options{"currency": "cad", "with_currency": true})
	expected := "$10.00 CAD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWhenJapanese(t *testing.T) {
	currency := New(10, Options{"currency": "jpy"})
	expected := "¥10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWhenJapaneseOverThousand(t *testing.T) {
	currency := New(1000, Options{"currency": "jpy"})
	expected := "¥1,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWhenJapaneseOverTenThousand(t *testing.T) {
	currency := New(10000, Options{"currency": "jpy"})
	expected := "¥10,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWhenJapaneseOverHundredThousand(t *testing.T) {
	currency := New(100000, Options{"currency": "jpy"})
	expected := "¥100,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWhenJapaneseOverMillion(t *testing.T) {
	currency := New(1000000, Options{"currency": "jpy"})
	expected := "¥1,000,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestNewWithoutCents(t *testing.T) {
	currency := New(10, Options{"with_cents": false})
	expected := "$10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestOverride(t *testing.T) {
	options := defaults()

	options = override(options, Options{"currency": "cad"})

	if options["currency"].(string) != "cad" {
		t.Error("Expected currency to be overriden to cad")
	}

	if options["with_currency"].(bool) != false {
		t.Error("Expected with_currency not to be overriden")
	}
}
