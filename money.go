package money

import (
	"fmt"
)

type Options map[string]interface{}

func addSymbol(result string, currency map[string]interface{}, options Options) string {
	var space string

	if options["with_symbol_space"].(bool) {
		space = " "
	}

	if currency["symbol_first"].(bool) {
		result = fmt.Sprintf("%s%s%s", currency["symbol"], space, result)
	} else {
		result = fmt.Sprintf("%s%s%s", result, space, currency["symbol"], space)
	}

	return result
}

func defaults() Options {
	return Options{
		"currency":                 "usd",
		"with_cents":               true,
		"with_currency":            false,
		"with_symbol":              true,
		"with_symbol_space":        false,
		"with_thousands_separator": true,
	}
}

func override(original, override Options) Options {
	for k, v := range override {
		if override[k] != nil {
			original[k] = v
		}
	}

	return original
}

func New(val float64, opts ...Options) (result string) {
	options := defaults()

	if len(opts) > 0 {
		options = override(options, opts[0])
	}

	currency := currencies[options["currency"].(string)]

	integer, fractional := splitValue(val)

	if options["with_thousands_separator"].(bool) {
		result = separateThousands(integer, currency["thousands_separator"].(string))
	} else {
		result = integer
	}

	if options["with_cents"].(bool) && currency["subunit"] != nil {
		result = fmt.Sprintf("%s%s%s", result, currency["decimal_mark"].(string), fractional)
	}

	if options["with_symbol"].(bool) {
		result = addSymbol(result, currency, options)
	}

	if options["with_currency"].(bool) {
		result = fmt.Sprintf("%s %s", result, currency["iso_code"])
	}

	return result
}
