package money

import (
	"testing"
)

func TestFormatWithDefaults(t *testing.T) {
	currency := Format(10)
	expected := "$10.00"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithCurrency(t *testing.T) {
	currency := Format(10, Options{"with_currency": true})
	expected := "$10.00 USD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithCurrencyWhenCanadian(t *testing.T) {
	currency := Format(10, Options{"currency": "CAD", "with_currency": true})
	expected := "$10.00 CAD"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapanese(t *testing.T) {
	currency := Format(10, Options{"currency": "JPY"})
	expected := "¥10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverThousand(t *testing.T) {
	currency := Format(1000, Options{"currency": "JPY"})
	expected := "¥1,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverTenThousand(t *testing.T) {
	currency := Format(10000, Options{"currency": "JPY"})
	expected := "¥10,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverHundredThousand(t *testing.T) {
	currency := Format(100000, Options{"currency": "JPY"})
	expected := "¥100,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWhenJapaneseOverMillion(t *testing.T) {
	currency := Format(1000000, Options{"currency": "JPY"})
	expected := "¥1,000,000"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestFormatWithoutCents(t *testing.T) {
	currency := Format(10, Options{"with_cents": false})
	expected := "$10"

	if currency != expected {
		t.Errorf("Expected %s but got %s", expected, currency)
	}
}

func TestAddSymbol(t *testing.T) {
	q := "10.00"

	if afn := addSymbol(q, currencies["AFN"], defaults()); afn != "10.00؋" {
		t.Errorf("Expected euro symbol to be placed after quantity, but got %s", afn)
	}

	if afn := addSymbol(q, currencies["AFN"], Options{"with_symbol_space": true}); afn != "10.00 ؋" {
		t.Errorf("Expected euro symbol to be placed after quantity and a space, but got %s", afn)
	}

	if usd := addSymbol(q, currencies["USD"], defaults()); usd != "$10.00" {
		t.Errorf("Expected dollar symbol to be placed before quantity, but got %s", usd)
	}

	if usd := addSymbol(q, currencies["USD"], Options{"with_symbol_space": true}); usd != "$ 10.00" {
		t.Errorf("Expected dollar symbol to be placed before quantity and a space, but got %s", usd)
	}
}

func TestSeparateThousands(t *testing.T) {
	values := map[string]string{
		"1":          "1",
		"12":         "12",
		"123":        "123",
		"1234":       "1,234",
		"12345":      "12,345",
		"69310":      "69,310",
		"123456":     "123,456",
		"1234567":    "1,234,567",
		"12345678":   "12,345,678",
		"123456789":  "123,456,789",
		"1234567891": "1,234,567,891",
	}

	for value, expected := range values {
		if v := separateThousands(value, ","); v != expected {
			t.Errorf("Expected %s to be %s", v, expected)
		}
	}
}
