package game

import (
	"github.com/MichaelThessel/spacee/app"
	"github.com/veandco/go-sdl2/sdl"
)

// Game holds the game state
type Game struct {
	a   *app.App
	p   *player
	pbl *bulletList // Player bullet list
	abl *bulletList // Alien bullet list
	ag  *alienGrid
	c   *Config
}

// Config holds game configuration
type Config struct {
	agc *alienGridConfig
	pc  *playerConfig
}

// New returns a new game
func New(a *app.App) (*Game, error) {
	var err error

	g := &Game{
		a:   a,
		pbl: &bulletList{},
		abl: &bulletList{},
	}
	g.initConfig()

	// Player
	g.p, err = newPlayer(a.GetRenderer(), g.c.pc)
	if err != nil {
		return nil, err
	}

	// Alien grid
	g.ag, err = newAlienGrid(a.GetRenderer(), g.c.agc)
	if err != nil {
		return nil, err
	}

	g.setup()

	return g, nil
}

// initConfig initalizes gthe game config
func (g *Game) initConfig() {
	g.c = &Config{
		agc: &alienGridConfig{
			rows:        5,
			cols:        10,
			marginRow:   20,
			marginCol:   20,
			returnPoint: 30,
			speed:       4,
			speedStep:   5,
			bulletSpeed: 15,
			fireRate:    0.14,
		},
		pc: &playerConfig{
			stepSize:    15,
			bulletSpeed: 15,
			lifes:       3,
		},
	}
}

// setup sets up the game
func (g *Game) setup() {
	g.a.RegisterKeyCallback(sdl.K_LEFT, func() { g.p.Move('l') })    // left
	g.a.RegisterKeyCallback(sdl.K_RIGHT, func() { g.p.Move('r') })   // right
	g.a.RegisterKeyCallback(sdl.K_SPACE, func() { g.p.Fire(g.pbl) }) // fire

	// Draw player
	g.a.RegisterRenderCallback(1, g.p.Draw)

	// Draw player & alien bullets
	g.a.RegisterRenderCallback(1, g.abl.Draw)
	g.a.RegisterRenderCallback(1, g.pbl.Draw)

	// Test if bullets have hit
	g.a.RegisterRenderCallback(1, func() { g.ag.testHit(g.pbl) })
	g.a.RegisterRenderCallback(1, func() { g.p.testHit(g.abl) })

	// Draw alien grid
	g.a.RegisterRenderCallback(1, g.ag.Draw)

	// Aliens fire
	g.a.RegisterRenderCallback(1, func() { g.ag.fire(g.abl) })
}
