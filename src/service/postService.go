package service

import (
	"dev-social-network/src/model"
	"dev-social-network/src/repository"
	"dev-social-network/src/utils"
	"encoding/json"
	"io"
)

func SavePost(body io.ReadCloser, userId string) (model.Post, error) {

	data, err := io.ReadAll(body)

	if err != nil {
		return model.Post{}, err
	}

	var post model.Post
	post.SetAuthor(userId)

	err = json.Unmarshal(data, &post)

	if err != nil {
		return model.Post{}, err
	}

	err = model.ValidatePost(post)

	if err != nil {
		return model.Post{}, err
	}

	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return model.Post{}, err
	}

	defer dbConnection.Close()

	post, err = repository.SavePost(dbConnection, post)

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func UpdatePost(body io.ReadCloser, userId string) (model.Post, error) {

	data, err := io.ReadAll(body)

	if err != nil {
		return model.Post{}, nil
	}

	var post model.Post
	post.SetAuthor(userId)

	err = json.Unmarshal(data, &post)

	if err != nil {
		return model.Post{}, err
	}

	err = model.ValidatePost(post)

	if err != nil {
		return model.Post{}, err
	}

	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return model.Post{}, err
	}

	defer dbConnection.Close()

	post, err = repository.UpdatePost(dbConnection, post)

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func DeletePost(body io.ReadCloser, userId string) error {
	data, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	var postId struct {
		Id string
	}

	err = json.Unmarshal(data, &postId)

	if err != nil {
		return err
	}

	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return err
	}

	defer dbConnection.Close()

	err = repository.DeletePost(dbConnection, postId.Id, userId)

	if err != nil {
		return err
	}

	return nil
}

func FindAllPosts(userId string, limit string, offset string) ([]model.Post, error) {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return []model.Post{}, err
	}

	defer dbConnection.Close()

	posts, err := repository.FindAllPosts(dbConnection, userId, limit, offset)

	if err != nil {
		return []model.Post{}, err
	}

	return posts, nil
}

func FindPost(postId string, userId string) (model.Post, error) {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return model.Post{}, err
	}

	defer dbConnection.Close()

	post, err := repository.FindPost(dbConnection, postId, userId)

	if err != nil {
		return model.Post{}, err
	}

	return post, nil
}

func LikePost(idPost string, idUser string) error {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return err
	}

	defer dbConnection.Close()

	err = repository.LikePost(dbConnection, idPost, idUser)

	if err != nil {
		return err
	}

	return nil

}

func DislikePost(postId string, userId string) error {
	dbConnection, err := utils.OpenConnection()

	if err != nil {
		return err
	}

	defer dbConnection.Close()

	err = repository.DislikePost(dbConnection, postId, userId)

	if err != nil {
		return err
	}

	return nil
}
