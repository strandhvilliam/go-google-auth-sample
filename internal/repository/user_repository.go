package repository

import (
	"fmt"
	"google-auth-go-v2/internal/infra"
	"google-auth-go-v2/internal/models"
)

func GetById(id string) (*models.User, error) {
	result := &models.User{}
	err := infra.GetDb().Get(result, "SELECT * FROM users WHERE id = ?", id)
	fmt.Println(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Create(user *models.User) error {
	_, err := infra.GetDb().Exec("INSERT INTO users (id, email, first_name, last_name, picture) VALUES (?, ?, ?, ?, ?)", user.Id, user.Email, user.FirstName, user.LastName, user.Picture)
	if err != nil {
		return err
	}
	return nil
}

func Update(user *models.User) error {
	_, err := infra.GetDb().Exec("UPDATE users SET email = ?, first_name = ?, last_name = ?, picture = ? WHERE id = ?", user.Email, user.FirstName, user.LastName, user.Picture, user.Id)
	if err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	_, err := infra.GetDb().Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
