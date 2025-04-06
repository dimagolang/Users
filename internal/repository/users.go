package repository

import (
	"Users/internal/models"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	db *pgx.Conn //
}

func NewUsersRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) InsertUserToDB(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO public.users (name, email, password)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(ctx, query, user.Name, user.Email, user.Password)
	if err != nil {
		return errors.Wrap(err, "failed to insert user into database")
	}

	fmt.Println("✅ Пользователь успешно добавлен в базу данных")
	return nil
}

func (s *UserRepo) UpdateUserInDB(ctx context.Context, email, name, password string) error {
	if email == "" || name == "" || password == "" {
		return errors.New("email, name и password не могут быть пустыми")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка хэширования пароля: %w", err)
	}

	updateQuery := `UPDATE users SET name = $1, password = $2 WHERE email = $3`
	result, err := s.db.Exec(ctx, updateQuery, name, string(hashedPassword), email)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении пользователя: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("пользователь не найден")
	}

	return nil // всё хорошо
}
