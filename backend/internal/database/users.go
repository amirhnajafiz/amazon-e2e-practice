package database

import (
	"context"
	"errors"

	"github.com/amirhnajafiz/aep/backend/pkg/models"
)

func (db Database) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	if count, err := db.conn.NewSelect().Model(&user).Where("username = ?", username).ScanAndCount(context.Background()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (db Database) CreateUser(user *models.User) (int64, error) {
	res, err := db.conn.NewInsert().Model(user).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
