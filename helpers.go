package money

import (
	"math"
	"strconv"
	"strings"
)

func separateThousands(value int64, separator string) string {
	v := strconv.FormatInt(value, 10)
	s := len(v) / 3
	m := int(math.Mod(float64(len(v)), 3))

	if m > 0 {
		s++
	}

	if s == 0 {
		return v
	}

	r := make([]string, s)

	for i := 0; i < len(r); i++ {
		if i == 0 && m > 0 {
			r[i] = v[i : i+m]
		} else {
			r[i] = v[i : i+3]
		}
	}

	return strings.Join(r, separator)
}
