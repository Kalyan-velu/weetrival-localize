package repositories

import (
	"context"
	"github.com/kalyan-velu/weetrival-localize/internal/db"
	"github.com/kalyan-velu/weetrival-localize/internal/models"
)

func CreateUser(ctx context.Context, user *models.User) error {
	_, err := db.DB.NewInsert().Model(user).Exec(ctx)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := db.DB.NewSelect().Model(&user).Where("email = ?").Limit(1).Scan(ctx)
	return &user, err
}
