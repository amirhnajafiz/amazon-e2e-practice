package database

import (
	"time"

	"github.com/amirhnajafiz/aep/backend/pkg/models"
)

func (db Database) GetAllUsers() ([]models.User, error) {
	var users []models.User

	if err := db.conn.Model(&users).Order("username ASC").Select(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db Database) GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	if err := db.conn.Model(&user).Where("username = ?", username).Select(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db Database) CreateUser(user *models.User) (int, error) {
	if res, err := db.conn.Model(user).Insert(); err != nil {
		return 0, err
	} else {
		return res.RowsAffected(), nil
	}
}

func (db Database) UpdateUser(user *models.User) (int, error) {
	if res, err := db.conn.Model(user).Where("username = ?", user.Username).Update(); err != nil {
		return 0, err
	} else {
		return res.RowsAffected(), nil
	}
}

func (db Database) DeleteUser(id string) (int, error) {
	if res, err := db.conn.Model(&models.User{}).Set("deleted_at", time.Now()).Where("username = ?", id).Update(); err != nil {
		return 0, err
	} else {
		return res.RowsAffected(), nil
	}
}
