package rng

// RNG : Random Number Generator
type RNG interface {
	Next() float64
}

type rng struct {
	a [4]int64
	b [4]int64
}

// NewRNG : create new RNG
func NewRNG(seed [6]int64) RNG {
	r := rng{
		[4]int64{0, seed[0], seed[1], seed[2]},
		[4]int64{0, seed[3], seed[4], seed[5]},
	}
	return RNG(&r)
}

func (this *rng) Next() float64 {
	this.a[0] = (1403580*this.a[2] - 810728*this.a[3]) % 4294967087
	this.b[0] = (527612*this.b[1] - 1370589*this.b[3]) % 4294944443
	var z = int64((this.a[0] - this.b[0]) % 4294967087)
	if z < 0 {
		z += 429467087
	}
	var u float64
	if z > 0 {
		u = float64(z) / 4294967087.0
	} else {
		u = 4294967087.0 / 4294967088.0
	}
	this.a[3] = this.a[2]
	this.a[2] = this.a[1]
	this.a[1] = this.a[0]
	this.b[3] = this.b[2]
	this.b[2] = this.b[1]
	this.b[1] = this.b[0]
	return u
}
