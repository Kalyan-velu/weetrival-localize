package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// AuthProvider represents the auth_providers table
type AuthProvider struct {
	bun.BaseModel `bun:"table:auth_providers"`

	ID         uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	UserID     uuid.UUID `bun:",notnull,type:uuid,rel:belongs-to,join:user_id=id"`
	Provider   string    `bun:",notnull,check:'provider IN (''google'', ''email_password'')'"`
	ProviderID string    `bun:",unique,notnull"` // Google ID or Email
	CreatedAt  time.Time `bun:",default:current_timestamp"`
}
