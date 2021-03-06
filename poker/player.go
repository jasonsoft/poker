package poker

import "time"

type Player struct {
	ID        int
	Name      string
	IsReal    bool
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
