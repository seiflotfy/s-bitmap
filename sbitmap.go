package sbitmap

import (
	"math"

	"github.com/dgryski/go-farm"
	"github.com/willf/bitset"
)

/*
Sbitmap represents a Self learning Bitmap structure
*/
type Sbitmap struct {
	V    *bitset.BitSet
	B    uint64
	L    uint64
	c    uint64
	d    uint64
	m    uint64
	nmax uint64
	r    float64
	err  float64
	err2 float64
	tt   float64
}

/*
New returns a Sbitmap with for the max cardinality nmax with an errorRate err
*/
func New(nmax uint64, err float64) *Sbitmap {
	s := &Sbitmap{}
	s.nmax = nmax
	s.err = err
	s.err2 = math.Pow(err, 2)
	mNum := math.Log(1 + 2*float64(nmax)*s.err2)
	mDenom := math.Log(1 + 2*s.err2*math.Pow(1-s.err2, -1))
	s.m = uint64(mNum / mDenom)
	s.V = bitset.New(uint(s.m))
	s.L = 0
	s.c = uint64(math.Log2(float64(s.m)))
	s.d = 64 - s.c
	s.r = 1 - 2*s.err2*math.Pow(1+s.err2, -1)
	s.tt = 0.0
	return s
}

/*
NewDefault returns a Sbitmap with for the max cardinality nmax math.MaxUint64 with an errorRate of 0.008
*/
func NewDefault() *Sbitmap {
	return New(math.MaxUint64, 0.008)
}

func (s *Sbitmap) getPk(k float64) float64 {
	m := float64(s.m)
	return m * math.Pow(m+1-k, -1) * (1 + s.err2) * math.Pow(s.r, k)
}

func (s *Sbitmap) getQ(l, pl float64) float64 {
	m := float64(s.m)
	return math.Pow(m, -1) * (m - l + 1) * pl
}

// Estimate returns the estimated cardinality of the Sbitmap
func (s *Sbitmap) Estimate() float64 {
	tl := 0.0
	for i := 0.0; i < float64(s.L)+1; i++ {
		q := math.Pow(s.getQ(i, s.getPk(i)), -1)
		tl += q
	}
	return tl
}

// Update adds item to the Sbitmap
func (s *Sbitmap) Update(item []byte) {
	h := farm.Hash64(item)
	j := h >> (s.d)
	if s.V.Test(uint(j)) == false {
		u := float64((h << s.c) >> s.c)
		pL1 := s.getPk(float64(s.L + 1))
		if u*math.Pow(2, -float64(s.d)) < pL1 {
			s.V.Set(uint(j))
			s.L++
		}
	}
}

// Contains returns membership of item in Sbitmap (can result in false positive)
func (s *Sbitmap) Contains(item []byte) bool {
	h := farm.Hash64(item)
	j := h >> (s.d)
	return s.V.Test(uint(j))
}
