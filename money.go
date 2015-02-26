package money

type Options map[string]interface{}

func defaults() Options {
	return Options{
		"currency":      "usd",
		"no_cents":      false,
		"with_currency": false,
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
