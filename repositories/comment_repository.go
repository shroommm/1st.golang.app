package repositories

import (
	"2nd.app/entities"
	"github.com/jmoiron/sqlx"
)

// CommentRepository provides actions to manipulate comments in the database.
type CommentRepository interface {
	Comment(ID int) (entities.Comment, error)
	Comments() ([]entities.Comment, error)
	Insert(comment entities.Comment) error
	Update(comment entities.Comment) error
	Delete(ID int) error
}

// commentRepo implements the CommentRepository interface.
type commentRepo struct {
	db *sqlx.DB
}

// NewCommentRepo instantiates commentRepo struct.
func NewCommentRepo(db *sqlx.DB) CommentRepository {
	return &commentRepo{
		db: db,
	}
}

// Comment fetches a comment from the database and return it.
func (r *commentRepo) Comment(ID int) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.Get(&comment, "SELECT * FROM comments WHERE id = $1", ID)
	return comment, err
}

// Comments fetches all the comments from the database and return them.
func (r *commentRepo) Comments() ([]entities.Comment, error) {
	var comments []entities.Comment
	err := r.db.Select(&comments, "SELECT * FROM comments")
	return comments, err
}

// Insert inserts a new comment into the database.
func (r *commentRepo) Insert(comment entities.Comment) error {
	_, err := r.db.NamedExec("INSERT INTO comments (user_id, post_id, description, status) VALUES (:user_id, :post_id, :description, :status)", comment)
	if err != nil {
		return err
	}

	return nil
}

// Update updates given comment in the database.
func (r *commentRepo) Update(comment entities.Comment) error {
	_, err := r.db.NamedExec("UPDATE comments SET user_id = :user_id, post_id = :post_id, description = :description, status = :status) WHERE id = :id", comment)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a comment from the database.
func (r *commentRepo) Delete(ID int) error {
	_, err := r.db.Exec("DELETE FROM comments WHERE id = $1", ID)
	if err != nil {
		return err
	}

	return nil
}
