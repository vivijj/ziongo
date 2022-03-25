package account

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/vivijj/ziongo/crypto/babyjub"
	"github.com/vivijj/ziongo/types/fr"
	"github.com/vivijj/ziongo/types/smt"
	"github.com/vivijj/ziongo/types/witness"
	"github.com/vivijj/ziongo/utils/hasher"
)

var (
	AccountHasher = hasher.NewPoseidonHasher(6)
)

// Account is zion network account
type Account struct {
	Address     common.Address
	PublicKey   babyjub.PublicKey
	Nonce       int
	Balances    map[int]*big.Int
	BalanceTree smt.SparseQuadMerkleTree
}

func (a *Account) IsEmpty() bool {
	return a.Address == common.Address{}
}

func (a *Account) BalanceRoot() fr.Repr {
	return a.BalanceTree.RootHash()
}

func (a *Account) VerifySignature(sig *babyjub.Signature, msg *big.Int) bool {
	return a.PublicKey.VerifyPoseidon(msg, sig)
}

func (a *Account) Hash() fr.Repr {
	address := fr.FromAddress(a.Address)
	publicKeyX := fr.FromBigInt(a.PublicKey.X)
	publicKeyY := fr.FromBigInt(a.PublicKey.Y)
	nonce := fr.FromInt(a.Nonce)
	root := a.BalanceRoot()

	return AccountHasher.HashFrRepr(
		[]fr.Repr{
			address,
			publicKeyX,
			publicKeyY,
			nonce,
			root,
		},
	)
}

// GetBalance return the token balance of this account, if token not exist, return 0
func (a *Account) GetBalance(tokenId int) *big.Int {
	if balance, ok := a.Balances[tokenId]; ok {
		return balance
	}
	return big.NewInt(0)
}

// UpdateBalance will update the Balances map and the BalanceTree in the same time.
func (a *Account) UpdateBalance(tokenId int, deltaBalance *big.Int) witness.BalanceUpdateWitness {
	// if this token not exist, insert with amount 0.
	if _, ok := a.Balances[tokenId]; !ok {
		a.Balances[tokenId] = big.NewInt(0)
	}

}
