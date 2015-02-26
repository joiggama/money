package money

func defaults() map[string]interface{} {
	return map[string]interface{}{
		"currency":      "usd",
		"no_cents":      false,
		"with_currency": false,
	}
}

func override(original, override map[string]interface{}) map[string]interface{} {
	for k, v := range override {
		if override[k] != nil {
			original[k] = v
		}
	}

	return original
}
