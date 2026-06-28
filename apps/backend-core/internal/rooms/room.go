package rooms

import (
	"errors"
	"sync"

	"github.com/amir138904/deborna-telegram-tma/apps/backend-core/internal/models"
)

/* =========================
   ROOM CORE (GAME + MONEY)
========================= */

type Room struct {
	mu sync.Mutex

	Players []*models.User

	EntryFee float64

	PlatformFeePercent float64

	ReferralReward float64

	Winner *models.User
}

/* =========================
   JOIN ROOM
========================= */

func (r *Room) Join(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.Balance < r.EntryFee {
		return errors.New("not enough balance")
	}

	// lock money
	user.Balance -= r.EntryFee
	user.PendingBalance += r.EntryFee

	r.Players = append(r.Players, user)

	return nil
}

/* =========================
   LEAVE ROOM (REFUND)
========================= */

func (r *Room) Leave(user *models.User) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.PendingBalance >= r.EntryFee {
		user.PendingBalance -= r.EntryFee
		user.Balance += r.EntryFee
	}
}

/* =========================
   START GAME
========================= */

func (r *Room) StartGame() {
	// engine will control real-time flow later
}

/* =========================
   SET WINNER
========================= */

func (r *Room) SetWinner(user *models.User) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.Winner = user
}

/* =========================
   END GAME + DISTRIBUTE MONEY
========================= */

func (r *Room) FinishGame() {
	r.mu.Lock()
	defer r.mu.Unlock()

	totalPool := float64(len(r.Players)) * r.EntryFee

	// 1. platform fee
	fee := (totalPool * r.PlatformFeePercent) / 100
	net := totalPool - fee

	// 2. referral reward (simple version: winner gets it handled here later)
	for _, p := range r.Players {
		p.PendingBalance = 0
	}

	// 3. payout
	if r.Winner != nil {
		r.Winner.Balance += net
	}
}


package rooms

import (
	"errors"
	"sync"

	"github.com/amir138904/deborna-telegram-tma/apps/backend-core/internal/models"
)

type Room struct {
	mu sync.Mutex

	Players []*models.User

	EntryFee float64

	PlatformFeePercent float64

	ReferralReward float64

	Winner *models.User
}