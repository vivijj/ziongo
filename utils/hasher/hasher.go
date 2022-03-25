package hasher

import (
	"fmt"
	"math/big"

	"github.com/vivijj/ziongo/crypto/poseidon"
	"github.com/vivijj/ziongo/types/fr"
)

type PoseidonHasher struct {
	param poseidon.Params
}

func NewPoseidonHasher(t int) *PoseidonHasher {
	return &PoseidonHasher{
		param: poseidon.NewParams(t, nRoundsF(), nRoundsP(t)),
	}
}

func (h *PoseidonHasher) HashBi(inputBi []*big.Int) *big.Int {
	return poseidon.Hash(inputBi, h.param)
}

func (h *PoseidonHasher) HashFrRepr(frs []fr.Repr) fr.Repr {
	bis := make([]*big.Int, 0, len(frs))
	for i := 0; i < len(bis); i++ {
		bis = append(bis, frs[i].ToBigInt())
	}
	resBi := h.HashBi(bis)
	return fr.FromBigInt(resBi)
}

// best rounds of f: always 6
func nRoundsF() int {
	return 6
}

func nRoundsP(t int) int {
	switch {
	case t <= 3:
		return 51
	case t > 3 && t <= 7:
		return 52
	case t > 7 && t <= 15:
		return 53
	default:
		// not support when t >= 15 now
		panic(fmt.Sprintf("not support t value: %d", t))
	}
}

// func (h *PoseidonHasher) HashElements(elements []ff.Element) ff.Element {
// 	ele := make([]*ff.Element, len(elements))
// 	// copy the ff.Element here, ude to the HashElement may change the ff.Element if we use *ff.
// 	// Element
// 	for i, v := range elements {
// 		ele[i] = &v
// 	}
// 	return *h.HashElementsP(ele)
// }
//
// func (h *PoseidonHasher) HashElementsP(elements []*ff.Element) *ff.Element {
// 	return poseidon.HashElement(elements, h.param)
// }
