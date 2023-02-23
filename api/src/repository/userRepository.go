package repository

import (
	"api/src/model"
	"api/src/utils"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

func CreateUser(db *sql.DB, user model.User) (model.User, error) {
	uuid := uuid.NewString()
	createdAt := time.Now().Format(time.RFC3339)

	hashedPass, erro := utils.HashPassWord(user.Password)

	if erro != nil {
		return model.User{}, erro
	}

	user.SetPassWord(string(hashedPass))

	statement, erro := db.Prepare(
		"INSERT INTO USER (ID,NAME,NICK_NAME,EMAIL,PASSWORD,CREATED_AT) VALUES(?,?,?,?,?,?)",
	)

	if erro != nil {
		return user, erro
	}

	result, erro := statement.Exec(uuid, user.Name, user.Nick_name, user.Email, user.Password, createdAt)

	if erro != nil {
		return user, erro
	}

	createdRows, erro := result.RowsAffected()

	if erro != nil {
		return user, erro
	}

	if createdRows > 0 {
		user.SetId(uuid)
		user.SetCreatedAt(createdAt)
		user.SetPassWord("")
	}

	defer statement.Close()

	return user, nil
}

func FindAll(db *sql.DB, limit string, offset string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	rows, erro := db.Query("SELECT ID, NAME, NICK_NAME, EMAIL FROM USER LIMIT ? OFFSET ?", limit, offset)

	if erro != nil {
		return []*model.User{}, erro
	}

	defer rows.Close()

	for rows.Next() {
		user := new(model.User)

		rows.Scan(&user.Id, &user.Name, &user.Nick_name, &user.Email)

		users = append(users, user)
	}

	return users, nil
}

func FindOne(db *sql.DB, id string) (model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT ID, NAME, NICK_NAME, EMAIL FROM USER WHERE ID = ?", id)

	if erro := row.Scan(&user.Id, &user.Name, &user.Nick_name, &user.Email); erro != nil {
		return user, erro
	}

	return user, nil
}

func UpdateUser(db *sql.DB, user model.User) (model.User, error) {
	statement, erro := db.Prepare(
		"UPDATE USER SET NAME = ?, NICK_NAME = ?, EMAIL = ? WHERE ID = ?",
	)

	if erro != nil {
		return user, erro
	}

	result, erro := statement.Exec(user.Name, user.Nick_name, user.Email, user.Id)

	if erro != nil {
		return user, erro
	}

	_, erro = result.RowsAffected()

	if erro != nil {
		return user, erro
	}

	defer statement.Close()

	return user, nil
}
func DeleteUser(db *sql.DB, id string) error {
	statement, erro := db.Prepare(
		"DELETE FROM USER WHERE ID = ?",
	)

	if erro != nil {
		return erro
	}

	result, erro := statement.Exec(id)

	if erro != nil {
		return erro
	}

	_, erro = result.RowsAffected()

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return nil
}

func LogIn(db *sql.DB, email string) (model.UserLogin, error) {
	var user model.UserLogin

	row := db.QueryRow("SELECT ID, EMAIL, PASSWORD FROM USER WHERE EMAIL = ?", email)

	if erro := row.Scan(&user.Id, &user.Email, &user.Password); erro != nil {
		return model.UserLogin{}, erro
	}

	return user, nil
}

func Follow(db *sql.DB, followingObj model.Following) error {
	statement, erro := db.Prepare(
		"INSERT INTO FOLLOWERS (ID_USER,ID_FOLLOWING) VALUES(?,?)",
	)

	if erro != nil {
		return erro
	}

	_, erro = statement.Exec(followingObj.UserId, followingObj.FollowerId)

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return nil
}

func UnFollow(db *sql.DB, idUser string, idFollowing string) error {
	statement, erro := db.Prepare(
		"DELETE FROM FOLLOWERS WHERE ID_USER = ? AND ID_FOLLOWING = ?",
	)

	if erro != nil {
		return erro
	}

	result, erro := statement.Exec(idUser, idFollowing)

	if erro != nil {
		return erro
	}

	_, erro = result.RowsAffected()

	if erro != nil {
		return erro
	}

	defer statement.Close()

	return nil
}

func FindFollowing(db *sql.DB, userId string, offset string, limit string) (interface{}, error) {
	type followers struct {
		Id       string
		Name     string
		NickName string
	}

	var returnedArr []followers

	rows, err := db.Query(`SELECT S.ID, S.NAME, S.NICK_NAME FROM FOLLOWERS F 
	INNER JOIN USER S ON F.ID_FOLLOWING = S.ID WHERE F.ID_USER = ? LIMIT ? OFFSET ?`, userId, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		follower := new(followers)

		rows.Scan(&follower.Id, &follower.Name, &follower.NickName)

		returnedArr = append(returnedArr, *follower)
	}

	return returnedArr, nil
}

func FindFollowers(db *sql.DB, userId string, offset string, limit string) (interface{}, error) {
	type followers struct {
		Id       string
		Name     string
		NickName string
	}

	var returnedArr []followers

	rows, err := db.Query(`SELECT S.ID, S.NAME, S.NICK_NAME FROM FOLLOWERS F 
	INNER JOIN USER S ON F.ID_USER = S.ID WHERE F.ID_FOLLOWING = ? LIMIT ? OFFSET ?`, userId, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		follower := new(followers)

		rows.Scan(&follower.Id, &follower.Name, &follower.NickName)

		returnedArr = append(returnedArr, *follower)
	}

	return returnedArr, nil
}
