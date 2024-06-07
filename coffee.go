package particles

import (
	"math"
	"math/rand"

    "github.com/theprimeagen/vim-with-me/pkg/assert"
)

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) rune {
    count := counts[row][col]
    if count < 3 {
        return ' '
    }
    if count < 6 {
        return '.'
    }
    if count < 9 {
        return ':'
    }
    if count < 12 {
        return '{'
    }

    return '}'
}

func reset(p *Particle, params *ParticleParams) {
	p.Lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.Speed = math.Floor(float64(params.MaxSpeed) * rand.Float64())

	maxX := math.Floor(float64(params.X / 2))
	x := math.Max(-maxX, math.Min(rand.NormFloat64(), maxX))
    p.X = x + maxX
	p.Y = 0
}

func nextPos(particle *Particle, deltaMS int64) {
	particle.Lifetime -= deltaMS
	if particle.Lifetime <= 0 {
		return
	}

	particle.Y += particle.Speed * (float64(deltaMS) / 1000.0) // Rise it straight up
}

func NewCoffee(width, height int) Coffee {
    assert.Assert(width % 2 == 0, "width of the particle system must be odd.")
	return Coffee{
		ParticleSystem: NewParticleSystem(
			ParticleParams{
				MaxLife:       7,
				MaxSpeed:      0.5,
				ParticleCount: 100,

				reset:        reset,
				ascii:        ascii, nextposition: nextPos,
			},
		),
	}
}
