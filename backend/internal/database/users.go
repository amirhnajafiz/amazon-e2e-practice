package database

import "github.com/amirhnajafiz/aep/backend/pkg/models"

func (db Database) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

	return users, nil
}

func (db Database) CreateUser(user *models.User) (*models.User, error) {
	// Here you would typically insert the user into the database
	// and return the created user object, possibly with an ID assigned by the database.
	return user, nil
}

func (db Database) UpdateUser(user *models.User) (*models.User, error) {
	// Here you would typically update the user in the database
	// and return the updated user object.
	return user, nil
}

func (db Database) DeleteUser(id string) error {
	// Here you would typically delete the user from the database
	// based on the provided ID.
	return nil
}
