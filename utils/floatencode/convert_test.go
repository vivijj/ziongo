package floatencode

import (
	"fmt"
	"math/big"
	"testing"
)

func TestRoundToFloatValue(t *testing.T) {
	tv, _ := new(big.Int).SetString("1000000000000000000000000000000000000000000", 10)
	rtv := RoundToFloatValue(tv, Float8Encoding)
	fmt.Println(rtv)
}
