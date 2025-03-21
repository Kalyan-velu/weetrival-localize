package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/kalyan-velu/weetrival-localize/internal/db"
	"github.com/kalyan-velu/weetrival-localize/internal/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	res, err := db.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return err
	}

	// Ensure at least one row was affected
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no user was inserted")
	}

	return nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.DB.NewSelect().Model(&user).Where("email = ?", email).Limit(1).Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Handle the case where no user is found
			return nil, nil // No user found is not an error
		}
		return nil, err // Other database errors
	}

	return &user, nil
}
