package particles

import (
	"math"
	"math/rand"
)

func assert(condition bool, message string) {
	if !condition {
		panic(message)
	}
}

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) string {
	count := counts[row][col]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "{"
	}
	return "}"
}

func reset(p *Particle, params *ParticleParams) {
	p.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64())) 
	p.Speed = params.MaxSpeed * rand.Float64()

	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(rand.NormFloat64(), maxX)) 
	p.X = x + maxX
	p.Y = 0
}

func nextPos(p *Particle, deltaMS int64) {
	p.Lifetime -= deltaMS
	if p.Lifetime <= 0 {
		return 
	}
	
	percent := (float64(deltaMS) / 1000.0)
	p.Y += p.Speed * percent
}

func NewCoffee(width, height int) Coffee {
	assert(width % 2 == 1, "width of the particle system MUST be odd")
	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7,
				MaxSpeed:      1,
				ParticleCount: 100,

				reset:        reset,
				ascii:        ascii,
				nextPosition: nextPos,

				X: width,
				Y: height,
			},
		),
	}
}