package floatencode

import "math/big"

type FloatEncoding struct {
	NumBitsExponent int64
	NumBitsMantissa int64
	ExponentBase    int64
}

var (
	Float28Encoding = FloatEncoding{5, 23, 10}
	Float24Encoding = FloatEncoding{5, 19, 10}
	Float16Encoding = FloatEncoding{5, 11, 10}
	Float12Encoding = FloatEncoding{5, 7, 10}
	Float8Encoding  = FloatEncoding{5, 3, 10}
)

func ToFloat(value *big.Int, floatEncoding FloatEncoding) int64 {
	ebase := big.NewInt(floatEncoding.ExponentBase)

	maxPower := big.NewInt(1<<floatEncoding.NumBitsExponent - 1)
	maxMantissa := big.NewInt(1<<floatEncoding.NumBitsMantissa - 1)
	maxExponent := new(big.Int).Exp(ebase, maxPower, nil)
	maxValue := new(big.Int).Mul(maxMantissa, maxExponent)

	if value.Cmp(maxValue) == 1 {
		panic("value too large")
	}

	exponent := 0
	r := new(big.Int).Div(value, maxMantissa)
	d := big.NewInt(1)
	for r.Cmp(ebase) >= 0 || value.Cmp(big.NewInt(0).Mul(d, maxMantissa)) == 1 {
		r = r.Div(r, ebase)
		exponent += 1
		d = d.Mul(d, ebase)
	}
	mantissa := big.NewInt(0).Div(value, d)

	if maxExponent.Cmp(big.NewInt(int64(exponent))) < 0 {
		panic("exponent too large.")
	}
	if mantissa.Cmp(maxMantissa) > 0 {
		panic("mantissa too large.")
	}
	f := int64(exponent<<floatEncoding.NumBitsMantissa) + mantissa.Int64()
	return f
}

func FromFloat(f int64, floatEncoding FloatEncoding) *big.Int {
	exponent := f >> floatEncoding.NumBitsMantissa
	mantissa := f & ((1 << floatEncoding.NumBitsMantissa) - 1)
	value := new(big.Int).Mul(
		big.NewInt(mantissa),
		new(big.Int).Exp(big.NewInt(int64(floatEncoding.ExponentBase)), big.NewInt(exponent), nil),
	)
	return value

}

func RoundToFloatValue(value *big.Int, encoding FloatEncoding) *big.Int {
	f := ToFloat(value, encoding)
	floatValue := FromFloat(f, encoding)
	return floatValue
}
