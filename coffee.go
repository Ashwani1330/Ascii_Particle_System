package particles

type Coffee struct {
	ParticleSystem
}

func ascii(row, col int, counts [][]int) rune {
	return '}'
}

func reset(p *Particle, params *ParticleParams) {
}

func nextPos(particle *Particle, deltaMS int64) {
}

func NewCoffee(width, height int) Coffee {

    return Coffee{
        ParticleSystem: NewParticleSystem(
            ParticleParams {
                MaxLife:       7,
                MaxSpeed:      0.5,
                ParticleCount: 100,

                reset: reset,
                ascii: ascii,
                nextposition: nextPos,
            },
        ),
    }
}
