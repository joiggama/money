/*
Package money is a library to deal with money and currency representation.
Inspired by ruby money library http://rubymoney.github.io/money.

Defaults

    Options{
      "currency":                 "USD",
      "with_cents":               true,
      "with_currency":            false,
      "with_symbol":              true,
      "with_symbol_space":        false,
      "with_thousands_separator": true,
    }

Usage

    Format(10)                                               // "$10.00"
    Format(10, Options{"currency": "EUR"})                   // "€10.00"
    Format(10, Options{"with_cents": false})                 // "$10"
    Format(10, Options{"with_currency:" true })              // "$10.00 USD"
    Format(10, Options{"with_symbol": false})                // "10.00"
    Format(10, Options{"with_symbol_space":true})            // "$ 10.00"
    Format(1000)                                             // "$1,000.00"
    Format(1000, Options{"with_thousands_separator": false}) // "$1000.00"
*/
package money

import (
	"fmt"
	"math"
	"strings"
)

// Format returns a formatted price string according to currency rules and options
func Format(val float64, opts ...Options) (result string) {
	options := defaults()

	if len(opts) > 0 {
		options = override(options, opts[0])
	}

	c := currencies[options["currency"].(string)]

	integer, fractional := splitValue(val)

	if options["with_thousands_separator"].(bool) {
		result = separateThousands(integer, c.ThousandsSeparator)
	} else {
		result = integer
	}

	if options["with_cents"].(bool) && c.SubUnit != "" {
		result = fmt.Sprintf("%s%s%s", result, c.DecimalMark, fractional)
	}

	if options["with_symbol"].(bool) {
		result = addSymbol(result, c, options)
	}

	if options["with_currency"].(bool) {
		result = fmt.Sprintf("%s %s", result, options["currency"].(string))
	}

	return result
}

func addSymbol(result string, c currency, options Options) string {
	var space string

	if options["with_symbol_space"].(bool) {
		space = " "
	}

	if c.SymbolFirst {
		result = fmt.Sprintf("%s%s%s", c.Symbol, space, result)
	} else {
		result = fmt.Sprintf("%s%s%s", result, space, c.Symbol)
	}

	return result
}

func separateThousands(value, separator string) string {
	chunks := len(value) / 3

	if chunks == 0 {
		return value
	}

	if partial := math.Mod(float64(len(value)), 3); partial > 0 {
		chunks++
	}

	result := make([]string, chunks)

	for i := chunks - 1; i >= 0; i-- {
		if i == 0 {
			result[i] = value
			break
		}

		chunk := value[len(value)-3:]
		value = strings.TrimSuffix(value, chunk)
		result[i] = chunk
	}

	return strings.Join(result, separator)
}

func splitValue(val float64) (integer, fractional string) {
	i, f := math.Modf(val)

	integer = fmt.Sprintf("%.0f", i)
	fractional = fmt.Sprintf("%.2f", f)[2:]

	return
}
