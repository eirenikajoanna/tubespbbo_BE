package service

import (
	"tubespbbo/hashing"
	"tubespbbo/modules/dto"
	"tubespbbo/modules/model"
	"tubespbbo/modules/repository"
)

func FindUser() (*[]model.User, error) {
	return repository.FindUser()
}

func FindOneUser(id int64) (*model.User, error) {
	return repository.FindOneUser(id)
}

func CreateUser(user *model.User) error {
	user.Password = hashing.HashAndSalt([]byte(user.Password))
	return repository.CreateUser(user)
}

func UpdateUser(updateDto *dto.UpdateUserDTO, id int64) (*model.User, error) {
	user, err := repository.FindOneUser(id)
	if err != nil {
		return nil, err
	}
	if updateDto.Address != "" {
		user.Address = updateDto.Address
	}
	if updateDto.Phone != "" {
		user.Phone = updateDto.Phone
	}
	if updateDto.Username != "" {
		user.Username = updateDto.Username
	}
	user, err = repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int64) (*model.User, error) {
	user, err := repository.FindOneUser(id)
	if err != nil {
		return nil, err
	}
	return repository.DeleteUser(user)
}
