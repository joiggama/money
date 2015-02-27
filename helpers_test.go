package money

import (
	"testing"
)

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
