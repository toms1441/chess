package order

import (
	"github.com/toms1441/chess-server/internal/board"
)

// [U]
type CredentialsModel struct {
	Token    string `json:"token"`
	PublicID string `json:"public_id"`
}

// [O]
type InviteModel struct {
	ID string `json:"id"`
}

// [U]
type GameModel struct {
	// which pieces are yours
	P1    bool         `json:"p1"`
	Board *board.Board `json:"board"`
}

// [O]
type MoveModel struct {
	ID  int8        `json:"id"`
	Dst board.Point `json:"dst"`
}

// [O] sent as response to http
type PossibleModel struct {
	ID     *int8         `json:"src,omitempty"`    // [C]
	Points *board.Points `json:"points,omitempty"` // [U]
}

// [U]
type TurnModel struct {
	P1 bool `json:"player"`
}

// [O]
type PromoteModel struct {
	ID   int   `json:"id"`
	Type uint8 `json:"type"`
}

// [U]
type PromotionModel PromoteModel /*struct {
	Type uint8       `json:"type"`
	Dst  board.Point `json:"dst"`
}*/

// [O]
type CastlingModel struct {
	Src int `json:"src"`
	Dst int `json:"dst"`
}

// [U]
type CheckmateModel TurnModel

// [O]
type DoneModel struct {
	P1 bool `json:"p1"`
}
