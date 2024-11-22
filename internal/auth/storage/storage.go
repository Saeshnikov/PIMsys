package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	auth_errors "pim-sys/internal/auth/errors"

	"github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

// SaveUser saves user to db.
func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte, name string, phone string) (int64, error) {
	const op = "storage.postgres.SaveUser"

	const isAdmin = false // will be fixed

	stmt, err := s.db.Prepare("INSERT INTO users(email, password, name, phone, isAdmin) VALUES($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		fmt.Println(err.Error())
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var id int64

	err = stmt.QueryRowContext(ctx, email, passHash, name, phone, isAdmin).Scan(&id)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 - unique_violation
			return 0, fmt.Errorf("%s: %w", op, auth_errors.ErrUserExists)
		}
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

// User returns user by email.
func (s *Storage) User(ctx context.Context, email string) (User, error) {
	const op = "storage.postgres.User"

	stmt, err := s.db.Prepare("SELECT id, email, password, isAdmin FROM users WHERE email = $1")
	if err != nil {
		return User{}, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email)

	var user User
	err = row.Scan(&user.ID, &user.Email, &user.PassHash, &user.IsAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, fmt.Errorf("%s: %w", op, auth_errors.ErrUserNotFound)
		}
		return User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// IsAdmin checks if the user is an admin.
func (s *Storage) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "storage.postgres.IsAdmin"

	stmt, err := s.db.Prepare("SELECT isAdmin FROM users WHERE id = $1")
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, userID)

	var isAdmin bool
	err = row.Scan(&isAdmin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, auth_errors.ErrUserNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
