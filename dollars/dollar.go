package dollars

import (
	"math"
	"strconv"
)

const precision = 100

type Dollar struct {
	units int64
}

func (d Dollar) GetValue() float64 {
	return float64(d.units) / precision
}

func FromInt(value int64) Dollar {
	return Dollar{
		units: value * precision,
	}
}

func FromFloat64(value float64) Dollar {
	return Dollar {
		units: int64(math.Ceil(value * precision)),
	}
}

func FromString(value string) (Dollar, error) {
	parsedValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return Dollar{}, err
	}
	return FromFloat64(parsedValue), nil
}
