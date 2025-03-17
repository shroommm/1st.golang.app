package repositories

import (
	"2nd.app/entities"
	"github.com/jmoiron/sqlx"
)

// PostRepository provides actions to manipulate posts in the database.
type PostRepository interface {
	Post(ID int) (entities.Post, error)
	Posts() ([]entities.Post, error)
	Insert(post entities.Post) error
	Update(post entities.Post) error
	Delete(ID int) error
}

// postRepo implements the PostRepository interface.
type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo instantiates postRepo struct.
func NewPostRepo(db *sqlx.DB) PostRepository {
	return &postRepo{
		db: db,
	}
}

// Post fetches a post from the database and return it.
func (r *postRepo) Post(ID int) (entities.Post, error) {
	var post entities.Post
	err := r.db.Get(&post, "SELECT * FROM posts WHERE id = $1", ID)
	return post, err
}

// Posts fetches all the posts from the database and return them.
func (r *postRepo) Posts() ([]entities.Post, error) {
	var posts []entities.Post
	err := r.db.Select(&posts, "SELECT * FROM posts")
	return posts, err
}

// Insert inserts a new post into the database.
func (r *postRepo) Insert(post entities.Post) error {
	_, err := r.db.NamedExec("INSERT INTO posts (user_id, title, description, status) VALUES (:user_id, :title, :description, :status)", post)
	if err != nil {
		return err
	}

	return nil
}

// Update updates given post in the database.
func (r *postRepo) Update(post entities.Post) error {
	_, err := r.db.NamedExec("UPDATE posts SET user_id = :user_id, title = :title, description = :description, status = :status) WHERE id = :id", post)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a post from the database.
func (r *postRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM posts WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
