package database

import (
	"context"
	"errors"
	"time"

	"github.com/amirhnajafiz/aep/backend/pkg/models"
)

func (db Database) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if count, err := db.conn.NewSelect().Model(&users).ScanAndCount(context.Background()); err != nil {
		return nil, err
	} else if count == 0 {
		return nil, errors.New("no users found")
	}

	return users, nil
}

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

func (db Database) UpdateUser(user *models.User) (int64, error) {
	res, err := db.conn.NewUpdate().Model(user).Where("username = ?", user.Username).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (db Database) DeleteUser(id string) (int64, error) {
	res, err := db.conn.NewUpdate().Model(&models.User{}).Set("deleted_at = ?", time.Now()).Where("username = ?", id).Exec(context.Background())
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
