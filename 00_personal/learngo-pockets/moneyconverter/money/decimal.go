package money

import (
	"strconv"
	"strings"
)

// Decimal can represent a floating-point number with a fixed precision.
// example: 1.52 = 152 * 10^(-2) and will be stored as {152, 2}
type Decimal struct {
	// subunits is the amount of subunits.
	// Multiply it by the precision to get th real value
	subunits  int64
	precision byte
}

func (d Decimal) ParseDecimal(str string) (Decimal, error) {
	before, after, ok := strings.Cut(str, ".")
	if !ok {
		before = str
		after = ""
	}

	num, err := strconv.ParseInt(before+after, 10, 64)
	if err != nil {
		return Decimal{}, nil
	}

	return Decimal{
		subunits:  num,
		precision: byte(len(after)),
	}, nil
}
