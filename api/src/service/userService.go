package service

import (
	"api/src/model"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"errors"
	"io"
	"strings"
)

func CreateUser(user io.ReadCloser) (model.User, error) {
	var userObj model.User
	data, erro := io.ReadAll(user)

	if erro != nil {
		return model.User{}, erro
	}

	if erro := json.Unmarshal(data, &userObj); erro != nil {
		return model.User{}, erro
	}

	erro = model.ValidateUser(userObj)
	if erro != nil {
		return model.User{}, erro
	}

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return model.User{}, erro
	}

	userObj, erro = repository.CreateUser(dbConnection, userObj)

	if erro != nil && strings.Contains(erro.Error(), "Duplicate entry") && strings.Contains(erro.Error(), "USER.EMAIL") {
		return userObj, errors.New("email already in use")
	} else if erro != nil {
		return userObj, erro
	}

	defer dbConnection.Close()

	return userObj, nil
}

func FindAll(limit string, offset string) ([]*model.User, error) {
	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return []*model.User{}, erro
	}

	users, erro := repository.FindAll(dbConnection, limit, offset)

	if erro != nil {
		return []*model.User{}, erro
	}

	defer dbConnection.Close()

	return users, nil
}

func FindOne(id string) (model.User, error) {

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return model.User{}, erro
	}

	user, erro := repository.FindOne(dbConnection, id)

	if erro != nil {
		return model.User{}, erro
	}

	return user, nil
}

func UpdateUser(user io.ReadCloser) (model.User, error) {
	var userObj model.User
	data, erro := io.ReadAll(user)

	if erro != nil {
		return model.User{}, erro
	}

	if erro := json.Unmarshal(data, &userObj); erro != nil {
		return model.User{}, erro
	}

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return model.User{}, erro
	}

	userObj, erro = repository.UpdateUser(dbConnection, userObj)

	if erro != nil {
		return userObj, erro
	}

	defer dbConnection.Close()

	return userObj, nil
}

func DeleteUser(id string) error {
	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return erro
	}

	erro = repository.DeleteUser(dbConnection, id)

	if erro != nil {
		return erro
	}

	defer dbConnection.Close()

	return nil
}

func Follow(body io.ReadCloser, user_id string) error {
	var idFollowing struct {
		Id string
	}
	data, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &idFollowing)

	if err != nil {
		return err
	}

	followIngObj := model.Following{UserId: user_id, FollowerId: idFollowing.Id}

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return erro
	}

	defer dbConnection.Close()

	err = repository.Follow(dbConnection, followIngObj)

	if err != nil && strings.Contains(err.Error(), "Duplicate entry") {
		return errors.New("already follow")
	}

	if err != nil {
		return err
	}

	return nil
}

func UnFollow(body io.ReadCloser, user_id string) error {
	var idFollowing struct {
		Id string
	}
	data, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &idFollowing)

	if err != nil {
		return err
	}

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		return erro
	}

	defer dbConnection.Close()

	err = repository.UnFollow(dbConnection, user_id, idFollowing.Id)

	if err != nil {
		return err
	}

	return nil
}

func FindFollowing(idUser string, offset string, limit string) (interface{}, error) {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer dbConnection.Close()

	results, err := repository.FindFollowing(dbConnection, idUser, offset, limit)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func FindFollowers(idUser string, offset string, limit string) (interface{}, error) {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return nil, err
	}

	defer dbConnection.Close()

	results, err := repository.FindFollowers(dbConnection, idUser, offset, limit)

	if err != nil {
		return nil, err
	}

	return results, nil
}
