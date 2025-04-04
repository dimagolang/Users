package repository

import (
	"Users/internal/models"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type UserRepo struct {
	db *pgx.Conn //
}

func NewUsersRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) InsertUserToDB(ctx context.Context, user *models.User) error {
	err := r.db.QueryRow(ctx,
		"INSERT INTO public.users (name, email, password) VALUES ($1, $2, $3) ",
		user.Name, user.Email, user.Password,
	)

	if err != nil {
		return errors.New("failed to insert user")
	}
	return nil
}

func (s *UserRepo) UpdateUserInDB(email, name, password string) error {
	// Проверяем, существует ли пользователь с указанным email
	var existingUser models.User
	query := `SELECT name, password FROM users WHERE email = $1`
	err := s.db.QueryRow(context.Background(), query, email).Scan(&existingUser.Name, &existingUser.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("пользователь не найден")
		}
		return err
	}

	// Обновляем данные пользователя
	updateQuery := `UPDATE users SET name = $1, password = $2 WHERE email = $3`
	_, err = s.db.Exec(context.Background(), updateQuery, name, password, email)
	if err != nil {
		return err
	}

	return nil
}
