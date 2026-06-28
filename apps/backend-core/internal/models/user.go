
package models

import (
	"time"

	"github.com/google/uuid"
)

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

type UserRole uint8

const (
	RolePlayer UserRole = iota + 1
	RoleModerator
	RoleAdmin
	RoleOwner
)

type UserStatus uint8

const (
	StatusOffline UserStatus = iota
	StatusOnline
	StatusPlaying
	StatusBanned
)

type CardLevel uint8

const (
	Bronze CardLevel = iota + 1
	Silver
	Gold
	Premium
)

type Statistics struct {
	TotalGames   uint64
	Wins         uint64
	Losses       uint64
	TotalPrize   float64
	HighestPrize float64
}

type User struct {
	ID         string
	TelegramID int64

	Username  string
	FirstName string
	LastName  string

	Language Language

	Role   UserRole
	Status UserStatus

	Level CardLevel

	Balance        float64
	PendingBalance float64 // 💰 پول قفل شده داخل بازی

	WalletAddress string

	Statistics Statistics

	IsPremium bool
	IsBanned  bool

	ReferredBy    string
	ReferralCount uint64

	HasPlayedFirstGame bool

	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(telegramID int64, username, first, last string) *User {
	now := time.Now()

	return &User{
		ID:         uuid.NewString(),
		TelegramID: telegramID,

		Username:  username,
		FirstName: first,
		LastName:  last,

		Language: LangEnglish,

		Role:   RolePlayer,
		Status: StatusOffline,

		Level: Bronze,

		Balance:        0,
		PendingBalance:  0,

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
