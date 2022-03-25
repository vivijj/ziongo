package smt

import (
	"testing"

	"github.com/vivijj/ziongo/utils/hasher"
)

func TestSmt(t *testing.T) {
	hashdd := hasher.NewPoseidonHasher(5)
	tr := New(16, "12314", *hashdd)

	for i := 10; i < 100000; i++ {
		tr.Update(i, "9999")
	}

}

func BenchmarkSmt(b *testing.B) {
	hashdd := hasher.NewPoseidonHasher(5)
	tr := New(16, "12314", *hashdd)
	for i := 0; i < b.N; i++ {
		for i := 10; i < 100; i++ {
			tr.Update(i, "9999")
		}
	}
}
