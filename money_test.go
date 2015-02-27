package money

import (
	"testing"
)

func TestNewWithDefaults(t *testing.T) {
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

func TestSeparateThousands(t *testing.T) {
	values := map[string]string{
		"1":         "1",
		"10":        "10",
		"100":       "100",
		"1000":      "1,000",
		"10000":     "10,000",
		"100000":    "100,000",
		"1000000":   "1,000,000",
		"10000000":  "10,000,000",
		"100000000": "100,000,000",
	}

	for value, expected := range values {
		if v := separateThousands(value, ","); v != expected {
			t.Error("Expected %s to be %s", v, expected)
		}
	}
}
