package models

import (
	"time"

	"github.com/google/uuid"
)

/* =========================
   LANGUAGE
========================= */

type Language string

const (
	LangEnglish Language = "en"
	LangPersian Language = "fa"
	LangSpanish Language = "es"
	LangFrench  Language = "fr"
	LangGerman  Language = "de"
	LangTurkish Language = "tr"
	LangArabic  Language = "ar"
)

/* =========================
   ROLE
========================= */

type UserRole uint8

const (
	RolePlayer UserRole = iota + 1
	RoleModerator
	RoleAdmin
	RoleOwner
)

/* =========================
   STATUS
========================= */

type UserStatus uint8

const (
	StatusOffline UserStatus = iota
	StatusOnline
	StatusPlaying
	StatusBanned
)

/* =========================
   CARD LEVEL
========================= */

type CardLevel uint8

const (
	Bronze CardLevel = iota + 1
	Silver
	Gold
	Premium
)

/* =========================
   STATISTICS
========================= */

type Statistics struct {
	TotalGames   uint64  `json:"total_games"`
	Wins         uint64  `json:"wins"`
	Losses       uint64  `json:"losses"`
	TotalPrize   float64 `json:"total_prize"`
	HighestPrize float64 `json:"highest_prize"`
}

/* =========================
   USER MODEL (UPDATED)
========================= */

type User struct {
	ID         string `json:"id" db:"id"`
	TelegramID int64  `json:"telegram_id" db:"telegram_id"`

	Username  string `json:"username" db:"username"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`

	Language Language `json:"language" db:"language"`

	Role   UserRole   `json:"role" db:"role"`
	Status UserStatus `json:"status" db:"status"`

	Level CardLevel `json:"level" db:"level"`

	Balance float64 `json:"balance" db:"balance"`

	WalletAddress string `json:"wallet_address" db:"wallet_address"`

	Statistics Statistics `json:"statistics" db:"statistics"`

	IsPremium bool `json:"is_premium" db:"is_premium"`
	IsBanned  bool `json:"is_banned" db:"is_banned"`

	/* =========================
	   NEW (REFERRAL SYSTEM)
	========================= */

	ReferredBy    string `json:"referred_by" db:"referred_by"`
	ReferralCount uint64 `json:"referral_count" db:"referral_count"`

	HasPlayedFirstGame bool `json:"has_played_first_game" db:"has_played_first_game"`

	/* =========================
	   TIME
	========================= */

	LastLogin time.Time `json:"last_login" db:"last_login"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

/* =========================
   CONSTRUCTOR
========================= */

func NewUser(telegramID int64, username, firstName, lastName string) *User {
	now := time.Now()

	return &User{
		ID:         uuid.NewString(),
		TelegramID: telegramID,

		Username:  username,
		FirstName: firstName,
		LastName:  lastName,

		Language: LangEnglish,

		Role:   RolePlayer,
		Status: StatusOffline,

		Level: Bronze,

		Balance: 0,

		Statistics: Statistics{},

		IsPremium: false,
		IsBanned:  false,

		ReferredBy:         "",
		ReferralCount:      0,
		HasPlayedFirstGame: false,

		LastLogin: now,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

/* =========================
   GAME METHODS
========================= */

func (u *User) Login() {
	u.Status = StatusOnline
	u.LastLogin = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Logout() {
	u.Status = StatusOffline
	u.UpdatedAt = time.Now()
}

func (u *User) StartGame() {
	u.Status = StatusPlaying

	if !u.HasPlayedFirstGame {
		u.HasPlayedFirstGame = true
	}

	u.UpdatedAt = time.Now()
}

func (u *User) EndGame(win bool, prize float64) {
	u.Status = StatusOnline

	u.Statistics.TotalGames++

	if win {
		u.Statistics.Wins++
		u.Statistics.TotalPrize += prize

		if prize > u.Statistics.HighestPrize {
			u.Statistics.HighestPrize = prize
		}
	} else {
		u.Statistics.Losses++
	}

	u.UpdatedAt = time.Now()
}

func (u *User) AddBalance(amount float64) {
	u.Balance += amount
	u.UpdatedAt = time.Now()
}

func (u *User) SubtractBalance(amount float64) bool {
	if u.Balance < amount {
		return false
	}

	u.Balance -= amount
	u.UpdatedAt = time.Now()
	return true
}

func (u *User) Promote(level CardLevel) {
	u.Level = level
	u.UpdatedAt = time.Now()
}
