package repositories

import (
	"2nd.app/entities"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	User(id string) (entities.User, error)
	Users() ([]entities.User, error)
	Insert(user entities.User) error
	Update(user entities.User) error
	Delete(id int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) User(id string) (entities.User, error) {
	var user entities.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userRepo) Users() ([]entities.User, error) {
	var users []entities.User
	err := r.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *userRepo) Insert(user entities.User) error {
	_, err := r.db.NamedExec("INSERT INTO users (firstname, lastname, email, status) VALUES (:firstname, :lastname, :email, :status)", user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Update(user entities.User) error {
	_, err := r.db.NamedExec("UPDATE users firstname = :firstname, lastname = :lastname, email = :email, status = :status WHERE id = :id", user)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
