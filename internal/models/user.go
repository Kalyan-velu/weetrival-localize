package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name         string    `bun:",notnull"`
	Email        string    `bun:",unique,notnull"`
	PasswordHash *string   `bun:",nullzero"` // Nullable for social login users
	Role         string    `bun:",notnull" sql:"CHECK (role IN ('admin', 'tourist','guide'))"`
	CreatedAt    time.Time `bun:",default:current_timestamp"`
}
