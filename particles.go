package particles

import (
	"math"
	"time"
)

type Particle struct {
	lifetime int64 // ms
	speed    float64

	x float64
	y float64
}

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	ParticleCount int

	X int
	Y int

	nextposition NextPositionFunc
	ascii        Ascii
	reset        Reset
}

type NextPositionFunc = func(particle *Particle, deltaMs int64)
type Ascii func(row, col int, count [][]int) rune // return type: rune
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
	}
}

func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.reset(p, &ps.ParticleParams)
	}
}

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.nextposition(p, delta)

		if p.y >= float64(ps.Y) || p.x >= float64(ps.X) {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func (ps *ParticleSystem) Display() [][]rune {
	counts := make([][]int, 0)

	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	for _, p := range ps.particles {
		row := int(math.Floor(p.y))
		col := int(math.Floor(p.x))

		counts[row][col]++
	}

	out := make([][]rune, 0)
	for r, row := range counts {
		outRow := make([]rune, 0)
		for c := range row {
			outRow = append(outRow, ps.ascii(r, c, counts))
		}

		out = append(out, outRow)
	}

	return out
}
