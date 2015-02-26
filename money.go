package money

func defaults() map[string]interface{} {
	return map[string]interface{}{
		"currency":      "usd",
		"no_cents":      false,
		"with_currency": false,
	}
}
