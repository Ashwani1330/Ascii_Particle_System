package particles

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	Lifetime int64 // ms
	Speed    float64

	X float64
	Y float64
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
type Ascii func(row, col int, count [][]int) string // return type: rune
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastTime int64
}

func NewParticleSystem(params ParticleParams) ParticleSystem {
    particles := make([]*Particle, 0)
    for i:= 0; i < params.ParticleCount; i++ {
        particles = append(particles, &Particle{})
    }
	return ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
        particles: particles,
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

		if p.Y >= float64(ps.Y) || p.X >= float64(ps.X) {
			ps.reset(p, &ps.ParticleParams)
		}
	}
}

func (ps *ParticleSystem) Display() string {
	counts := make([][]int, 0)

	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

    fmt.Printf("particles : %d %d %+v\n", ps.Y, ps.X, counts)

	for _, p := range ps.particles {
		row := int(math.Floor(p.Y))
		col := int(math.Floor(p.X))

		counts[row][col]++
	}

	out := make([][]string, 0)
	for r, row := range counts {
		outRow := make([]string, 0)
		for c := range row {
			outRow = append(outRow, ps.ascii(r, c, counts))
		}

        out = append(out, outRow)
	}

    slices.Reverse(out)
    outStr := make([]string, 0)
    for _, row := range out {
        outStr = append(outStr, strings.Join(row, "\n"))
    }

	return strings.Join(outStr, "\n")
}
