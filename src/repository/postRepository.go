package repository

import (
	"database/sql"
	"dev-social-network/src/model"
	"time"

	"github.com/google/uuid"
)

func SavePost(db *sql.DB, post model.Post) (model.Post, error) {
	uuid := uuid.NewString()
	createdAt := time.Now().UTC().Format("2006-01-02 03:04:05")

	statement, err := db.Prepare(`INSERT INTO POSTS (ID, TITLE, CONTENT, AUTHOR, CREATED_AT) VALUES 
	(?,?,?,?,?)`)

	if err != nil {
		return model.Post{}, err
	}

	defer statement.Close()

	result, err := statement.Exec(uuid, post.Title, post.Content, post.Author, createdAt)

	if err != nil {
		return model.Post{}, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return model.Post{}, err
	}

	if rowsAffected > 0 {
		post.SetId(uuid)
		post.SetCreatedAt(createdAt)

		row, err := db.Query("SELECT LIKES, DISLIKES FROM POSTS WHERE ID = ?", uuid)

		if err != nil {
			return model.Post{}, err
		}

		for row.Next() {
			row.Scan(&post.Likes, &post.Dislikes)
		}

		defer row.Close()
	} else {
		return model.Post{}, nil
	}

	return post, nil
}

func UpdatePost(db *sql.DB, post model.Post) (model.Post, error) {
	updatedAt := time.Now().UTC().Format("2006-01-02 03:04:05")

	statement, err := db.Prepare("UPDATE POSTS SET TITLE = ?, CONTENT = ?, UPDATED_AT = ? WHERE ID = ?")

	if err != nil {
		return model.Post{}, err
	}

	result, err := statement.Exec(post.Title, post.Content, updatedAt, post.Id)

	if err != nil {
		return model.Post{}, err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return model.Post{}, err
	}

	if rowAffected > 0 {
		row, err := db.Query(`SELECT LIKES, DISLIKES, CREATED_AT, UPDATED_AT FROM 
		POSTS WHERE ID = ?`, post.Id)

		if err != nil {
			return model.Post{}, err
		}

		for row.Next() {
			row.Scan(&post.Likes, &post.Dislikes, &post.CreatedAt, &post.UpdatedAt)
		}
	}

	return post, nil
}

func DeletePost(db *sql.DB, postId string, userId string) error {

	statement, err := db.Prepare("DELETE FROM POSTS WHERE ID = ? AND AUTHOR = ?")

	if err != nil {
		return err
	}

	result, err := statement.Exec(postId, userId)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

func FindAllPosts(db *sql.DB, userId string, limit string, offset string) ([]model.Post, error) {
	rows, err := db.Query(`SELECT 
	ID, TITLE, CONTENT, LIKES, DISLIKES, CREATED_AT, UPDATED_AT 
	FROM POSTS WHERE AUTHOR = ? LIMIT ? OFFSET ?`, userId, limit, offset)

	if err != nil {
		return []model.Post{}, err
	}

	var posts []model.Post

	for rows.Next() {
		post := new(model.Post)
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.Likes, &post.Dislikes, &post.CreatedAt, &post.UpdatedAt)

		posts = append(posts, *post)
	}

	return posts, nil
}

func FindPost(db *sql.DB, postId string, userId string) (model.Post, error) {
	row, err := db.Query(`SELECT ID, TITLE, CONTENT, AUTHOR, LIKES, 
	DISLIKES, CREATED_AT, UPDATED_AT FROM POSTS WHERE ID = ? AND AUTHOR = ?`, postId, userId)

	if err != nil {
		return model.Post{}, err
	}

	var post model.Post

	for row.Next() {
		row.Scan(&post.Id, &post.Title, &post.Content, &post.Author, &post.Likes, &post.Dislikes, &post.CreatedAt,
			&post.UpdatedAt)
	}

	return post, nil
}
