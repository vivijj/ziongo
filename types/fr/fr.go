package fr

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

// Repr is in the decimal representation
// e.g. "231231231231231312312312312312312313123"
// shouldn't be the hex string.
type Repr string

func FromBigInt(bi *big.Int) Repr {
	return Repr(bi.String())
}

func FromAddress(addr common.Address) Repr {
	addrBytes := addr.Bytes()
	addrBi := new(big.Int).SetBytes(addrBytes)
	return Repr(addrBi.String())
}

func FromInt(i int) Repr {
	return Repr(strconv.Itoa(i))
}

func (fr Repr) ToBigInt() *big.Int {
	res, _ := new(big.Int).SetString(string(fr), 10)
	return res
}

func FrsToBigInt(frs []Repr) []*big.Int {
	bis := make([]*big.Int, 0, len(frs))

	for i := range frs {
		bi, _ := new(big.Int).SetString(string(frs[i]), 10)
		bis = append(bis, bi)
	}
	return bis
}
