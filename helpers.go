package money

import (
	"fmt"
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

func splitValue(val float64) (integer, fractional string) {
	i, f := math.Modf(val)

	integer = fmt.Sprintf("%.0f", i)
	fractional = fmt.Sprintf("%.2f", f)[2:]

	return
}
