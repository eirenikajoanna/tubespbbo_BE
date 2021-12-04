package repository

import (
	"errors"
	"tubespbbo/db"
	"tubespbbo/modules/model"
)

func FindUser() (*[]model.User, error) {
	var users []model.User
	result := db.Orm.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return &users, nil
}

func FindOneUser(id int64) (*model.User, error) {
	var user model.User
	result := db.Orm.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("no user found")
	}

	return &user, nil
}

func FindUserByUsername(username string) (*model.User, error) {
	var user model.User
	_ = db.Orm.Where("username = ?", username).First(&user)
	return &user, nil
}

func CreateUser(user *model.User) error {
	result := db.Orm.Create(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(user *model.User) (*model.User, error) {
	result := db.Orm.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	user, _ = FindOneUser(user.Id)
	return user, nil
}

func DeleteUser(user *model.User) (*model.User, error) {
	result := db.Orm.Delete(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
