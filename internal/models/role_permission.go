package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// RolePermission represents the roles_permissions table
type RolePermission struct {
	bun.BaseModel `bun:"table:roles_permissions"`

	ID         uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Role       string    `bun:",notnull,check:'role IN (''admin'', ''tourist'', ''guide'')'"`
	Permission string    `bun:",notnull"`
	CreatedAt  time.Time `bun:",default:current_timestamp"`
}
