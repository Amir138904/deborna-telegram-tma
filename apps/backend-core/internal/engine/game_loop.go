package engine

import (
	"time"
	"sync"

	"github.com/amir138904/deborna-telegram-tma/apps/backend-core/internal/rooms"
	"github.com/amir138904/deborna-telegram-tma/apps/backend-core/internal/models"
)

type GameLoop struct {
	mu sync.Mutex

	Room *rooms.Room

	IsRunning bool

	Duration time.Duration
}

func NewGameLoop(room *rooms.Room, duration time.Duration) *GameLoop {
	return &GameLoop{
		Room:     room,
		Duration: duration,
	}
}

func (g *GameLoop) Start() {
	g.mu.Lock()
	if g.IsRunning {
		g.mu.Unlock()
		return
	}
	g.IsRunning = true
	g.mu.Unlock()

	go g.run()
}

func (g *GameLoop) run() {
	time.Sleep(5 * time.Second)

	g.Room.StartGame()

	timer := time.NewTimer(g.Duration)
	<-timer.C

	g.endGame()
}

func (g *GameLoop) endGame() {
	g.mu.Lock()
	defer g.mu.Unlock()

	if len(g.Room.Players) == 0 {
		g.IsRunning = false
		return
	}

	// فعلاً ساده: اولین نفر برنده
	winner := g.Room.Players[0]

	g.Room.SetWinner(winner)
	g.Room.FinishGame()

	g.IsRunning = false
}