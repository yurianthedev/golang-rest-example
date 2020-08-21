package postgres

import (
	"database/sql"
	"log"

	"github.com/yurianxdev/rest-example/api/model"
)

type Repo struct {
	DB *sql.DB
}

func (ur *Repo) ListUsers() ([]model.User, error) {
	rows, err := ur.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	var users []model.User
	var userCounter int
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Address)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
		userCounter++
	}

	log.Printf("Got %d rows\n", userCounter)
	return users, nil
}

func (ur *Repo) CreateUser(user model.User) (model.User, error) {
	result, err := ur.DB.Exec("INSERT INTO users (name, email, phone, address) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Phone, user.Address)
	if err != nil {
		return model.User{}, err
	}

	rowsAffected, _ := result.RowsAffected()
	idCreated, _ := result.LastInsertId()
	log.Printf("%d rows inserted\n", rowsAffected)

	user.Id = uint(idCreated)

	return user, nil
}
