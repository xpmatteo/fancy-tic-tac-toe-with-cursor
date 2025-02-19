package randomai

import "math/rand"

type RealRNG struct{}

func NewRealRNG() *RealRNG {
	return &RealRNG{}
}

func (r *RealRNG) NextValue(n int) int {
	return rand.Intn(n)
}
