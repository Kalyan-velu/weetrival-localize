package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// RefreshToken represents the refresh_tokens table
type RefreshToken struct {
	bun.BaseModel `bun:"table:refresh_tokens"`

	ID        uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	UserID    uuid.UUID `bun:",notnull,type:uuid,rel:belongs-to,join:user_id=id"`
	TokenHash string    `bun:",unique,notnull"` // Hashed token for security
	ExpiresAt time.Time `bun:",notnull"`
	CreatedAt time.Time `bun:",default:current_timestamp"`
}
