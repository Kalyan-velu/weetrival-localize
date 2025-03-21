package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()"`
	Name         string    `bun:",notnull"`
	Email        string    `bun:",unique,notnull"`
	PasswordHash *string   `bun:",nullzero" json:"-"` // Nullable for social login users
	Role         string    `bun:",notnull" sql:"CHECK (role IN ('admin', 'tourist','guide'))"`
	CreatedAt    time.Time `bun:",default:current_timestamp"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts a User model to a UserResponse DTO.
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

// ValidateRole checks  for user role and throws error.
func (u *User) ValidateRole() error {
	validRoles := map[string]bool{"admin": true, "tourist": true, "guide": true}
	if !validRoles[u.Role] {
		return fmt.Errorf("invalid role: %s", u.Role)
	}
	return nil
}
