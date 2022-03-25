package smt

import (
	"github.com/vivijj/ziongo/types/fr"
	"github.com/vivijj/ziongo/utils/hasher"
)

const Nary = 4

type SparseQuadMerkleTree struct {
	Depth  int
	Hasher hasher.PoseidonHasher
	Cache  map[fr.Repr][Nary]fr.Repr
	Root   fr.Repr
}

func New(depth int, defaultLeafHash fr.Repr, hasher hasher.PoseidonHasher) *SparseQuadMerkleTree {
	h := defaultLeafHash
	cache := make(map[fr.Repr][Nary]fr.Repr)
	for i := 0; i < depth; i++ {
		currentLayerH := [Nary]fr.Repr{h, h, h, h}
		newh := hasher.HashFrRepr(currentLayerH[:])
		cache[newh] = currentLayerH
		h = newh
	}
	return &SparseQuadMerkleTree{
		Depth:  depth,
		Hasher: hasher,
		Cache:  cache,
		Root:   h,
	}
}

func (s *SparseQuadMerkleTree) RootHash() fr.Repr {
	return s.Root
}

func (s *SparseQuadMerkleTree) GetHash(leafIndex int) fr.Repr {
	v := s.Root
	lookup := leafIndex

	for i := 0; i < s.Depth; i++ {
		childIndex := (lookup >> (2 * (s.Depth - 1))) % Nary
		v = s.Cache[v][childIndex]
		lookup <<= 2
	}
	return v
}

// Update will execute the `upsert` since the tree is always "full", so when we first insert,
// we just update the default value.
func (s *SparseQuadMerkleTree) Update(index int, itemHash fr.Repr) {
	v := s.Root
	lookupRef := index
	updateRef := index
	var sideNodes [][Nary]fr.Repr

	// lookup the path in the tree of the target node, record the path node hash.
	for i := 0; i < s.Depth; i++ {
		children := s.Cache[v]
		sideNodes = append(sideNodes, children)
		childIndex := (lookupRef >> 2 * (s.Depth - 1)) % Nary
		v = children[childIndex]
		lookupRef <<= 2
	}

	v = itemHash

	// update the merkle tree bottom up
	for i := 0; i < s.Depth; i++ {
		childIndex := updateRef % Nary
		var leaves [Nary]fr.Repr
		for c := 0; c < Nary; c++ {
			if c != childIndex {
				leaves[c] = sideNodes[s.Depth-1-i][c]
			} else {
				leaves[c] = v
			}
		}
		newV := s.Hasher.HashFrRepr(leaves[:])
		s.Cache[newV] = leaves
		updateRef >>= 2
		v = newV
	}
	s.Root = v
}

// MerklePath create the proof of existence for a certain element of the tree.
func (s *SparseQuadMerkleTree) MerklePath(index int) []fr.Repr {
	v := s.Root
	lookupRef := index
	var sideNodes [][]fr.Repr
	for i := 0; i < s.Depth; i++ {
		childIndex := (lookupRef >> (2 * (s.Depth - 1))) % Nary
		for c := 0; c < Nary; c++ {
			if c != childIndex {
				sideNodes[s.Depth-1-i] = append(sideNodes[s.Depth-1-i], s.Cache[v][c])
			}
		}
		v = s.Cache[v][childIndex]
		lookupRef *= Nary
	}

	// due to that len(sideNodes[0]) should always be 3
	numLevelNode := len(sideNodes[0])
	merkleProof := make([]fr.Repr, 0, len(sideNodes)*numLevelNode)
	for i := range sideNodes {
		for j := range sideNodes[i] {
			merkleProof = append(merkleProof, sideNodes[i][j])
		}
	}
	if !s.VerifyProof(merkleProof, index, v) {
		panic("not valid proof.")
	}
	return merkleProof
}

// VerifyProof verify the given merkle proof and verify if the calculate root is same with
// current one
func (s *SparseQuadMerkleTree) VerifyProof(
	merkleProof []fr.Repr,
	index int,
	itemHash fr.Repr,
) bool {
	lookupRef := index
	v := itemHash
	proofIndex := 0
	for i := 0; i < s.Depth; i++ {
		var input []fr.Repr
		for c := 0; c < Nary; c++ {
			if lookupRef%Nary == c {
				input = append(input, v)
			} else {
				input = append(input, merkleProof[proofIndex])
				proofIndex += 1
			}
		}
		newV := s.Hasher.HashFrRepr(input)
		lookupRef /= Nary
		v = newV
	}
	return s.Root == v
}
