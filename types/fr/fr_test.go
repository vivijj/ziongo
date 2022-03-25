package fr

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestFr(t *testing.T) {
	frr := Repr("1000000000000000")
	fmt.Println(frr.ToBigInt())

}

func TestFromAddress(t *testing.T) {
	addr := common.HexToAddress("0x2a500A5e1950aea40C22d8885C8DC3c02e99b3E2")
	fmt.Println(FromAddress(addr))
}

func BenchmarkFromAddress2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addr := common.HexToAddress("0x2a500A5e1950aea40C22d8885C8DC3c02e99b3E2")
		FromAddress(addr)
	}
}
