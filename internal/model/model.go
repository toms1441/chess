package model

import (
	"github.com/toms1441/chess-server/internal/board"
)

type OAuth2Config struct {
	ClientID     string `validate:"required" mapstructure:"client_id"`
	ClientSecret string `validate:"required" mapstructure:"client_secret"`
	Redirect     string `validate:"required" mapstructure:"redirect"`
}

// Profile is a generic struct used to get information about the user.
type Profile struct {
	ID       string `json:"id" validate:"required"`
	Picture  string `json:"picture" validate:"required"`
	Username string `json:"username" validate:"required"`
	Platform string `json:"platform" validate:"required"`
}

func (p Profile) Valid() bool {
	return len(p.ID) > 0 && len(p.Picture) > 0 && len(p.Username) > 0
}

// GetInviteID returns the invite id from a specificed token, this function's purpose is to have a standard way of retrieving ids...
func (p Profile) GetInviteID() string {
	return p.ID + "_" + p.Platform
}

// Possible is a struct indicating which moves are legal.
type Possible struct {
	ID     int8          `json:"id,omitempty" validate:"required"` // [C]
	Points *board.Points `json:"points,omitempty"`                 // [U]
}

// Watchable is a game that could be spectated
type Watchable struct {
	P1  Profile      `json:"p1"`
	P2  Profile      `json:"p2"`
	Brd *board.Board `json:"brd,omitempty"`
}

// Generic is a struct used for generic stuff. If a model *only* uses ID, then it should use generic.
type Generic struct {
	ID string `json:"id"`
}
