package model

import "errors"

type User struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick_name string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (u *User) SetId(id string) {
	u.Id = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetNickName(nickName string) {
	u.Nick_name = nickName
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetPassWord(password string) {
	u.Password = password
}

func (u *User) SetCreatedAt(createdAt string) {
	u.CreatedAt = createdAt
}

func ValidateUser(user User) error {

	if len(user.Name) <= 0 {
		return errors.New("empty 'Name' field")
	} else if len(user.Nick_name) <= 0 {
		return errors.New("empty 'Nick name' field")
	} else if len(user.Email) <= 0 {
		return errors.New("empty 'email' field")
	} else if len(user.Password) <= 0 {
		return errors.New("empty 'password' field")
	}

	return nil
}
