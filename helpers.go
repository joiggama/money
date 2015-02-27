package money

import (
	"math"
	"strings"
)

func separateThousands(value, separator string) string {
	s := len(value) / 3
	m := int(math.Mod(float64(len(value)), 3))

	if m > 0 {
		s++
	}

	if s == 0 {
		return value
	}

	r := make([]string, s)

	for i := 0; i < len(r); i++ {
		if i == 0 && m > 0 {
			r[i] = value[i : i+m]
		} else {
			r[i] = value[i : i+3]
		}
	}

	return strings.Join(r, separator)
}
