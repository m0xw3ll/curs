package src

import (
	"math"
    "math/big"
)

const (
	lns int = 20
)

type Rand struct {
	rnd  *big.Int
	iter *big.Int
	base *big.Int
	l    *big.Int
}

func NewGen(Base, L int) *Rand {
	out := new(Rand)
	out.base = base
	out.iter = 1
	out.L = L
	return out
}

func (rnd *Rand) slice() []int {
	base := rnd.base
	out := make([]int, lns)
	temp_base := 1
	for i := lns - 1; i >= 0; i-- {
		out[i] = temp_base
		temp_base *= base
	}
	return out
}

func (rnd *Rand) bslice() []int {
	out := make([]int, lns)
	temp := rnd.rnd
	base_t := rnd.slice()
	for i := 0; i < lns; i++ {
		temp_s := temp / base_t[i]
		out[i] = temp_s
		temp %= base_t[i]
	}
	return out
}

func (rnd *Rand) Next() float64 {
	rnd.rnd = int(math.Floor(float64(rnd.iter) * math.Sqrt(float64(rnd.L)*float64(rnd.iter))))
	//rnd.rnd = rnd.iter * int(math.Floor(math.Sqrt(float64(rnd.base))))
	slice := rnd.bslice()
	out := float64(0)
	for i := 0; i < lns; i++ {
		out += float64(slice[lns-1-i]) * math.Pow(float64(rnd.base), float64(-i-1))
	}
	rnd.iter++
	return out
}
