package account

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddress(t *testing.T) {
	a := common.Address{}
	fmt.Println(a)
	b := common.Address{}

	fmt.Println(a == b)
}
