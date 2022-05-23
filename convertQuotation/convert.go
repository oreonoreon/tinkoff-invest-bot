package convertQuotation

import (
	"fmt"
	"math"
)

func Convert(units int64, nano int32) float64 {

	return float64(units) + float64(nano)/math.Pow(10, float64(len(fmt.Sprint(nano))))
}
